package player

import (
	"Muth/config"
	"bytes"
	"fmt"
	"github.com/gammazero/deque"
	"github.com/lonelyevil/kook"
	"os"
	"os/exec"
	"strconv"
	"sync"
	"syscall"
	"time"
)

type Music struct {
	File string
}

type ThreadManager struct {
	cmd    *exec.Cmd
	cancel func(*exec.Cmd)
	mu     sync.Mutex
	onExit func()
}

type Player struct {
	NowPlaying  *Music
	Queue       *deque.Deque[*Music]
	ControlChan chan int
	Mutex       sync.Mutex
	Manager     *ThreadManager
	Ctx         *kook.KmarkdownMessageContext
	Channel     int
	Volume      int
}

var MusicPlayer *Player

func (music *Music) Copy() *Music {

	return &Music{
		File: music.File,
	}
}

func (tm *ThreadManager) StartThread(onExit func(), player *Player) error {
	logger := config.Logger

	tm.mu.Lock()
	defer tm.mu.Unlock()

	if tm.cmd != nil && tm.cancel != nil {
		tm.cancel(tm.cmd)
		tm.cmd.Wait()
	}

	cmd := exec.Command(config.Config.KOOKVoice,
		"-c",
		strconv.Itoa(player.Channel),
		"-i",
		player.NowPlaying.File,
		"-t",
		config.Config.BotToken,
		"-af",
		fmt.Sprintf("loudnorm=i=-16,volume=%ddB", player.Volume),
	)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
	}

	tm.cancel = func(cmd *exec.Cmd) {
		syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
	}

	tm.cmd = cmd
	tm.onExit = onExit

	var outputBuffer bytes.Buffer
	cmd.Stdout = &outputBuffer
	cmd.Stderr = &outputBuffer

	go func() {
		if err := cmd.Run(); err != nil {
			logger.Err(err).Msgf("Error when running thread: %s", outputBuffer.String())
		}

		time.Sleep(5 * time.Second)

		if tm.onExit != nil {
			tm.onExit()
		}
	}()

	return nil
}

func (tm *ThreadManager) StopThread() {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	if tm.cmd != nil && tm.cancel != nil {
		tm.cancel(tm.cmd)
		tm.cmd.Wait()
		tm.cmd = nil
		tm.cancel = nil
	}
}

func NewPlayer() *Player {
	player := new(Player)
	logger := config.Logger

	err := os.RemoveAll("./assets/voice")
	if err != nil {
		logger.Err(err).Msg("Error when removing voice folder")
	}
	err = os.Mkdir("./assets/voice", 0755)
	if err != nil {
		logger.Err(err).Msg("Error when creating voice folder")
	}

	player.Queue = deque.New[*Music]()
	player.ControlChan = make(chan int)
	player.Mutex = sync.Mutex{}
	player.Manager = &ThreadManager{}
	player.Channel = config.Config.VoiceChannel[0].ID
	player.Volume = config.Config.DefaultVolume
	go player.Worker()

	return player
}

func (player *Player) AddMusic(music *Music) {
	player.Mutex.Lock()
	defer player.Mutex.Unlock()

	player.Queue.PushBack(music)
}

func (player *Player) SetCtx(ctx *kook.KmarkdownMessageContext) {
	player.Ctx = ctx
}

func (player *Player) SetVolume(volume int) {
	player.Mutex.Lock()
	defer player.Mutex.Unlock()

	player.Volume = volume

	if player.NowPlaying == nil {
		return
	}

	player.Queue.PushFront(player.NowPlaying.Copy())
	player.Manager.StopThread()
}

func (player *Player) SetChannel(channel int) {
	player.Mutex.Lock()
	defer player.Mutex.Unlock()

	player.Channel = channel

	if player.NowPlaying == nil {
		return
	}

	player.Queue.PushFront(player.NowPlaying.Copy())
	player.Manager.StopThread()
}

func (player *Player) Worker() {
	logger := config.Logger

	for {

		if player.NowPlaying == nil && player.Queue.Len() != 0 {
			player.Mutex.Lock()

			player.NowPlaying = player.Queue.PopFront()
			logger.Info().Msg("Now playing: " + player.NowPlaying.File)
			logger.Info().Msg("Queue length: " + strconv.Itoa(player.Queue.Len()))

			player.Mutex.Unlock()

			logger.Info().Msg("Starting thread: " + player.NowPlaying.File)
			player.Manager.StartThread(func() {
				player.NowPlaying = nil
				logger.Info().Msg("Set NowPlaying to nil")
			}, player)
		}
	}
}

func (player *Player) SendMsg(content string) {
	_, _ = player.Ctx.Session.MessageCreate(&kook.MessageCreate{
		MessageCreateBase: kook.MessageCreateBase{
			TargetID: player.Ctx.Common.TargetID,
			Content:  content,
		},
	})
}

package frontend

import (
	"log"
	"os"

	"github.com/farm-er/tchat/network"
	"github.com/farm-er/tchat/user"
	"github.com/gdamore/tcell/v2"
)






type Screen struct {

	screen tcell.Screen

	w int

	h int

	cursorX int

	cursorY int

	message string

	user *user.User

	port int
}




func NewScreen( user *user.User, port int) ( *Screen, error){
	
	screen, r := tcell.NewScreen()

	if r != nil {
		return nil, r
	}

	// other fields are initialized inside the after
	return &Screen{
		user: user,
		screen: screen,
		port: port,
	}, nil

}


func (s *Screen) Start( mainUser *user.User) {

	s.screen.Init()

	s.w, s.h = s.screen.Size()


	s.drawMainLayout( mainUser)

	// place the cursor 
	s.screen.ShowCursor( s.cursorX, s.cursorY)

	s.screen.Show()

	recChan := make(chan struct{})
	senChan := make(chan struct{})

	for {

		w, h := s.screen.Size()

		if w != s.w || h != s.h {
			s.w = w
			s.h = h
			s.drawMainLayout( mainUser)
		}

		s.drawMainLayout( mainUser)

		for i, r := range s.message {
			s.screen.SetContent( int(s.w/3)+3+i, s.h-2, r, nil, tcell.Style{})
		}

		s.screen.ShowCursor(s.cursorX, s.cursorY)

		s.screen.Show()


		switch ev := s.screen.PollEvent().(type) {
		case *tcell.EventKey:
		
			switch ev.Key() {
			case tcell.KeyUp:
				mainUser.ShiftFocusP()
			case tcell.KeyDown:
				mainUser.ShiftFocusN()
			case tcell.KeyCtrlA:
				// TODO: start sending signals 

				go func () {

					if r := network.SendSignals( senChan, s.port); r != nil {
						log.Fatal("Error sending signals ", r)
					}

				}()

			case tcell.KeyCtrlQ:
				// TODO: start receiving signals 
				go func(){

					if r := network.ReceiveSignals(recChan, s.port); r != nil {
						log.Fatal("Error receiving signals ", r)
					}

				}()

			case tcell.KeyCtrlC:

				senChan <- struct{}{}

			case tcell.KeyCtrlX:

				recChan <- struct{}{}

			case tcell.KeyEsc:
				s.screen.Fini()
				os.Exit(0)
			case tcell.KeyEnter:
				// TODO: send text

				// moving the cursor 
				s.cursorX = s.cursorX - len(s.message)
				// cleaning the ui
				r := make( []rune, len(s.message)-1)
				for i:= range r {
					r[i] = ' '
				}
				s.screen.SetContent( s.cursorX, s.cursorY, ' ', r, tcell.Style{})
				// clearing the input 
				s.message = ""
			case tcell.KeyBackspace, tcell.KeyBackspace2, tcell.KeyDelete:
				if len(s.message) > 0 {
					s.deleteText()
				} 
			case tcell.KeyRune:
				s.typeText(ev.Rune())
			}
		}

	}

}


func (s *Screen) drawMainLayout( user *user.User) {

	s.screen.Clear()

	// draw separator
	for i:= 0;i<s.h;i++ {
		s.screen.SetContent( int(s.w/3), i, '|', nil, tcell.Style{})
	}

	// TODO: draw conversations boxes 

	reachedHeight := 3

	limiter := reachedHeight + 6

	for i, mem  := range user.Members {

		style := tcell.StyleDefault 

		if i == user.MemFocus {
			// TODO: change style for the focused conversation 
			style = style.Foreground(tcell.ColorBlue)
		}

		// conversation box's top border 
		for i:= 1; i<int(s.w/3)-1; i++ {
			s.screen.SetContent( i, reachedHeight, '-', nil, style)
			s.screen.SetContent( i, limiter, '-', nil, style)
		}

		// conversation box's username 
		for i, v := range mem.GetUsername() {
			s.screen.SetContent( i+5, (limiter - reachedHeight) / 2 + reachedHeight, v, nil, style)
		}
		
		// conversation box's right border 
		for i:= reachedHeight +1 ; i<limiter; i++ {
			s.screen.SetContent( int(s.w/3)-1, i, '|', nil, style)
			s.screen.SetContent( 0, i, '|', nil, style)
		}

		reachedHeight += 7
		limiter = reachedHeight + 6

	}

	// side borders of input box 
	s.screen.SetContent( int(s.w/3)+1, s.h-2, '|', nil, tcell.Style{})
	s.screen.SetContent( s.w-1, s.h-2, '|', nil, tcell.Style{})

	// top and bottom borders of input box
	for i:=int(s.w/3)+2;i<s.w-1;i++ {
		s.screen.SetContent( i, s.h-1, '-', nil, tcell.Style{})
		s.screen.SetContent( i, s.h-3, '-', nil, tcell.Style{})
	}
	
	// store cursor's position
	s.cursorX = int(s.w/3)+3+len(s.message)
	s.cursorY = s.h-2

}


func (s *Screen) typeText(r rune) {
	
	s.message = s.message + string(r)
	
	s.cursorX += 1

	s.screen.ShowCursor( s.cursorX, s.cursorY)
}

func (s *Screen) deleteText() {

	s.message = s.message[:len(s.message)-1]
	
	s.screen.SetContent( int(s.w/3)+3+len(s.message), s.h-2, ' ', nil, tcell.Style{})

	s.cursorX -= 1 

	s.screen.ShowCursor( s.cursorX, s.cursorY)

}


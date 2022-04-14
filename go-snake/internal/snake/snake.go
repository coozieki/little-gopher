package snake

type Snake interface {
	Move(dir Direction)
	PushBlock()
	GetBlocks() []Block
}

type snake struct {
	blocks     []Block
	currentDir Direction
	blocksMap  map[float64]map[float64]bool
}

func NewSnake() *snake {
	return &snake{
		blocks:     []Block{{X: 0, Y: 0, Direction: DirectionRight}},
		currentDir: DirectionRight,
		blocksMap: map[float64]map[float64]bool{
			0: {0: true},
		},
	}
}

func (s *snake) PushBlock() {
	firstBlock := s.blocks[len(s.blocks)-1]

	var x, y = firstBlock.X, firstBlock.Y
	switch s.currentDir {
	case DirectionRight:
		x++
	case DirectionDown:
		y++
	case DirectionUp:
		y--
	case DirectionLeft:
		x--
	default:
		return
	}
	s.blocks = append(s.blocks, Block{X: x, Y: y, Direction: s.currentDir})
	_, ok := s.blocksMap[x]
	if !ok {
		s.blocksMap[x] = map[float64]bool{}
	}
	s.blocksMap[x][y] = true
}

func (s *snake) GetBlocks() []Block {
	return s.blocks
}

func (s *snake) Move(dir Direction) {
	temp := s.blocks[len(s.blocks)-1]
	temp.Direction = dir
	switch dir {
	case DirectionRight:
		if s.currentDir == DirectionLeft {
			return
		}
		temp.X++
	case DirectionDown:
		if s.currentDir == DirectionUp {
			return
		}
		temp.Y++
	case DirectionLeft:
		if s.currentDir == DirectionRight {
			return
		}
		temp.X--
	case DirectionUp:
		if s.currentDir == DirectionDown {
			return
		}
		temp.Y--
	}
	if s.blocksMap[temp.X][temp.Y] {
		return
	}
	s.currentDir = dir
	firstBlock := s.blocks[0]
	s.blocksMap[firstBlock.X][firstBlock.Y] = false
	s.blocks = s.blocks[1:]
	s.blocks = append(s.blocks, temp)
	_, ok := s.blocksMap[temp.X]
	if !ok {
		s.blocksMap[temp.X] = map[float64]bool{}
	}
	s.blocksMap[temp.X][temp.Y] = true
}

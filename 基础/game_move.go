package main
import (
	"fmt"
	"math"
)

type Vec2 struct {
	X,Y float32
}

func (v Vec2) Add(other Vec2) Vec2{
	return Vec2{
		v.X+other.X,
		v.Y+other.Y,
	}
}

func (v Vec2) Sub(other Vec2) Vec2{
	return Vec2{
		v.X-other.X,
		v.Y-other.Y,
	}
}

func (v Vec2) Scale(s float32) Vec2{
	return Vec2{
		v.X * s,
		v.Y * s,
	}
}

func (v Vec2) Distance(other Vec2) float32{
	v1 := v.Sub(other)
	return float32(math.Sqrt(float64(v1.X*v1.X+v1.Y*v1.Y)))
}

//单位化
func (v Vec2) Normalize() Vec2{
	mag := v.X*v.X+v.Y*v.Y
	if mag > 0 {
		oneOverMag := 1/float32(math.Sqrt(float64(mag)))
		return Vec2{
			v.X*oneOverMag,
			v.Y*oneOverMag,
		}
	}

	return Vec2{0,0}
}



type Player struct{
	currPos Vec2 //当前位置
	targetPos Vec2 //目标位置
	speed float32 //移动速度
}

func (p *Player) MoveTo(v Vec2){
	p.targetPos = v
}

func (p *Player) Pos() Vec2{
	return p.currPos
}

//是否到达
func (p *Player) IsArrived() bool {
	return p.currPos.Distance(p.targetPos) < p.speed
}

func (p *Player) Update(){
	if !p.IsArrived(){
		dir := p.targetPos.Sub(p.currPos).Normalize()
        newPos := p.currPos.Add(dir.Scale(p.speed))
		p.currPos = newPos
	}
}


func NewPlayer (speed float32) *Player{
	return &Player{
		speed:speed,
	}
}

func main() {
	p := NewPlayer(0.5)
	p.MoveTo(Vec2{3,1})

	for !p.IsArrived(){
		p.Update()
		fmt.Println(p.Pos())
	}
}

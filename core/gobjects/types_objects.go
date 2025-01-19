package gobjects

type Type int

const (
    Player Type = iota
	Enemy
	Bullet
    Tile
)


type ObjectType interface {
    Type() Type
}


func(t Type) Type() Type {
    return t
}
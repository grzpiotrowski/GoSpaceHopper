package main

type Monster struct {
	Entity
}

func makeMonster() (*Monster, error) {
	e := &Monster{
		Entity: NewEntity(Vec2f{0.4, 0.4}),
	}

	e.Movement.Speed = Vec2f{40, 40}

	img, _, err := NewImageFromFile("data/images/monster_stand.png")
	if err != nil {
		return nil, err
	}

	e.Graphics.Sprite = NewSprite(img)

	return e, nil
}

func (g *Game) updateMonster() {
	t := g.monster.Transform
	m := g.monster.Movement

	t.Position.X += m.Velocity.X * dt
	t.Position.Y += m.Velocity.Y * dt
}

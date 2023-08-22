package main

type Monster struct {
	Entity
}

func makeMonster() (*Monster, error) {
	e := &Monster{
		Entity: NewEntity(Vec2f{0.4, 0.4}),
	}

	e.Movement.Speed = Vec2f{40, 40}
	e.Transform.Position.X = 350

	img, _, err := NewImageFromFile("data/images/monster_stand.png")
	if err != nil {
		return nil, err
	}

	e.Graphics.Sprite = NewSprite(img)

	return e, nil
}

func (monster *Monster) Scroll(m MovementComponent) {
	monster.Transform.Position.X -= m.Velocity.X * dt
}

func (g *Game) updateMonster() {

	t := g.monster.Transform
	m := g.monster.Movement

	t.Position.Y += m.Velocity.Y * dt

	t.Position.Y += gravity * dt
	g.monster.Scroll(*g.hero.Movement)
	g.monster.TerrainBlock = g.monster.Entity.getBlockUnder(g.t)
	g.monster.Entity.update()
}

package main

type Hero struct {
	Entity
}

func makeHero() (*Hero, error) {
	e := &Hero{
		Entity: NewEntity(Vec2f{0.5, 0.5}),
	}

	e.Movement.Speed = Vec2f{120, 120}
	e.Transform.Position.X = 175

	img, _, err := NewImageFromFile("data/images/hero_stand.png")
	if err != nil {
		return nil, err
	}

	e.Graphics.Sprite = NewSprite(img)

	return e, nil
}

func (g *Game) updateHero() {
	t := g.hero.Transform
	m := g.hero.Movement

	t.Position.Y += m.Velocity.Y * dt

	t.Position.Y += gravity * dt
	g.hero.TerrainBlock = g.hero.Entity.getBlockUnder(g.t)
	g.hero.Entity.update()

}

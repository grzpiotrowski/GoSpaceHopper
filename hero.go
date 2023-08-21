package main

type Hero struct {
	Entity
}

func makeHero() (*Hero, error) {
	e := &Hero{
		Entity: NewEntity(Vec2f{0.5, 0.5}),
	}

	e.Movement.Speed = Vec2f{120, 120}

	img, _, err := NewImageFromFile("data/images/hero_stand.png")
	if err != nil {
		return nil, err
	}

	e.Graphics.Sprite = NewSprite(img)

	return e, nil
}

func (h *Hero) getBlockUnder(t *Terrain) *TerrainBlock {
	xHero := h.getAABB().X + h.getAABB().W/2

	for _, block := range t.Blocks {
		if xHero >= block.xBegin && xHero < block.xEnd() {
			return &block
		}
	}

	return nil

}

func (g *Game) updateHero() {
	t := g.hero.Transform
	m := g.hero.Movement

	t.Position.X += m.Velocity.X * dt
	t.Position.Y += m.Velocity.Y * dt

	t.Position.Y += gravity * dt
	tbUnder := g.hero.getBlockUnder(g.t)
	g.hero.Entity.update(*tbUnder)

}

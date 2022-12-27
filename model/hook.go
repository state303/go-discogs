package model

import (
	"github.com/state303/go-discogs/src/cache"
	"gorm.io/gorm"
)

func (g *Genre) AfterCreate(_ *gorm.DB) (err error) {
	cache.GenreCache.Store(g.Name, g.ID)
	return
}
func (g *Genre) AfterUpdate(_ *gorm.DB) (err error) {
	cache.GenreCache.Store(g.Name, g.ID)
	return
}
func (s *Style) AfterCreate(_ *gorm.DB) (err error) {
	cache.StyleCache.Store(s.Name, s.ID)
	return
}
func (s *Style) AfterUpdate(_ *gorm.DB) (err error) {
	cache.StyleCache.Store(s.Name, s.ID)
	return
}

func (m *Master) AfterCreate(_ *gorm.DB) (err error) {
	cache.MasterIDCache.Store(m.ID, struct{}{})
	return
}

func (m *Master) AfterUpdate(_ *gorm.DB) (err error) {
	cache.MasterIDCache.Store(m.ID, struct{}{})
	return
}

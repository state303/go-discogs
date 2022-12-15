package batch

import "github.com/knadh/koanf"

func hasArtist(k *koanf.Koanf) bool {
	return k.Bool("artists")
}

func hasLabel(k *koanf.Koanf) bool {
	return k.Bool("labels")
}

func hasMaster(k *koanf.Koanf) bool {
	return k.Bool("masters")
}

func hasRelease(k *koanf.Koanf) bool {
	return k.Bool("releases")
}

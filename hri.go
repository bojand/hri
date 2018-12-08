// Package hri contains utility functions for generating human readable string ids.
package hri

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	adjectives = []string{"able", "afraid", "ancient", "angry", "average", "bad", "big", "bitter",
		"black", "blue", "brave", "breezy", "bright", "brown", "busy", "calm", "careful", "chatty",
		"cheap", "chilly", "clean", "clear", "clever", "close", "cloudy", "cold", "comfortable",
		"cool", "cowardly", "cuddly", "curly", "curvy", "cute", "dangerous", "dapper", "dark",
		"dead", "deep", "difficult", "dirty", "dry", "dull", "early", "empty", "evil", "exciting",
		"expensive", "fair", "famous", "far", "fast", "fat", "fine", "flat", "fluffy", "foolish",
		"free", "fresh", "friendly", "full", "funny", "fuzzy", "gentle", "giant", "good", "great",
		"green", "grumpy", "happy", "hard", "healthy", "heavy", "helpless", "high", "honest",
		"horrible", "hot", "hungry", "important", "interesting", "itchy", "jolly", "kind", "large",
		"late", "lazy", "light", "little", "long", "loud", "lovely", "low", "lucky", "massive",
		"mean", "mighty", "modern", "moody", "narrow", "nasty", "near", "neat", "nervous", "new",
		"nice", "noisy", "odd", "old", "orange", "ordinary", "perfect", "pink", "plastic", "polite",
		"popular", "pretty", "proud", "purple", "quick", "quiet", "rare", "red", "rich", "rotten",
		"rude", "sad", "safe", "salty", "selfish", "serious", "shaggy", "share", "sharp", "short",
		"shy", "silent", "silly", "slimy", "slippery", "slow", "small", "smart", "smooth", "soft",
		"sour", "spicy", "splendid", "spotty", "stale", "strange", "strong", "stupid", "sweet",
		"swift", "tall", "tame", "tasty", "tender", "terrible", "thick", "thin", "thirsty", "tidy",
		"tiny", "tough", "tricky", "ugly", "unlucky", "useful", "warm", "weak", "wet", "white",
		"whole", "wicked", "windy", "wise", "witty", "wonderful", "yellow", "young"}

	nouns = []string{"ape", "baboon", "badger", "bat", "bear", "bird", "bobcat", "bulldog",
		"bullfrog", "cat", "catfish", "cheetah", "chicken", "chipmunk", "cobra", "cougar", "cow",
		"crab", "deer", "dingo", "dodo", "dog", "dolphin", "donkey", "dragon", "dragonfly", "duck",
		"eagle", "earwig", "eel", "elephant", "emu", "falcon", "fireant", "firefox", "fish", "fly",
		"fox", "frog", "gecko", "goat", "goose", "grasshopper", "horse", "hound", "husky", "impala",
		"insect", "jellyfish", "kangaroo", "ladybug", "liger", "lion", "lionfish", "lizard",
		"mayfly", "mole", "monkey", "moose", "moth", "mouse", "mule", "newt", "octopus", "otter",
		"owl", "panda", "panther", "parrot", "penguin", "pig", "puma", "pug", "quail", "rabbit",
		"rat", "rattlesnake", "robin", "seahorse", "sheep", "shrimp", "skunk", "sloth", "snail",
		"snake", "squid", "starfish", "stingray", "swan", "termite", "tiger", "treefrog",
		"turkey", "turtle", "vampirebat", "walrus", "warthog", "wasp", "wolverine", "wombat",
		"yak", "zebra"}

	source = rand.NewSource(time.Now().UnixNano())
	random = rand.New(source)
)

// Random returns a random human readable string id.
func Random() string {
	adj := adjectives[random.Intn(len(adjectives))]
	noun := nouns[random.Intn(len(nouns))]
	num := random.Intn(100)
	return fmt.Sprintf("%s-%s-%d", adj, noun, num)
}

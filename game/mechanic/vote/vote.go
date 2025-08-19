package vote

import (
	"github.com/google/uuid"
	"github.com/thronesmc/game/game/utils/maputils"
)

var (
	// votes stores the vote count per pollID and option
	votes = maputils.NewMap[string, map[string]int]()

	// userVote stores which option each player voted for in each pollID
	userVote = maputils.NewMap[string, map[uuid.UUID]string]()
)

// InitPoll initializes a poll with empty options.
func InitPoll(pollID string, options []string) {
	if _, ok := votes.Load(pollID); !ok {
		optionMap := make(map[string]int)
		for _, opt := range options {
			optionMap[opt] = 0
		}
		votes.Store(pollID, optionMap)
		userVote.Store(pollID, make(map[uuid.UUID]string))
	}
}

// CastVote registers a player's vote for a poll.
func CastVote(pollID string, playerUUID uuid.UUID, option string) {
	if pollVotes, ok := votes.Load(pollID); ok {
		if pollUserVotes, ok := userVote.Load(pollID); ok {
			// If the player already voted, decrease the previous option count
			if oldOption, voted := pollUserVotes[playerUUID]; voted {
				pollVotes[oldOption]--
			}
			// Register the new vote
			pollVotes[option]++
			pollUserVotes[playerUUID] = option
		}
	}
}

// RemoveVote removes a player's vote from a poll.
func RemoveVote(pollID string, playerUUID uuid.UUID) {
	if pollVotes, ok := votes.Load(pollID); ok {
		if pollUserVotes, ok := userVote.Load(pollID); ok {
			if option, voted := pollUserVotes[playerUUID]; voted {
				// Decrease the option count
				pollVotes[option]--
				// Remove the player's vote record
				delete(pollUserVotes, playerUUID)
			}
		}
	}
}

// GetPollResults returns the current vote counts for a poll.
func GetPollResults(pollID string) map[string]int {
	if pollVotes, ok := votes.Load(pollID); ok {
		// Create a copy to prevent external modification
		result := make(map[string]int)
		for k, v := range pollVotes {
			result[k] = v
		}
		return result
	}
	return nil
}

// GetTopOption returns the option with the highest votes in a poll.
// Returns empty string if no votes have been cast.
func GetTopOption(pollID string) string {
	if pollVotes, ok := votes.Load(pollID); ok {
		var topOption string
		var maxVotes int
		for option, count := range pollVotes {
			if count > maxVotes || topOption == "" {
				topOption = option
				maxVotes = count
			}
		}
		return topOption
	}
	return ""
}

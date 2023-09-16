package govapi

import "time"

type GetAccountByNameParams struct {
	Name  string
	Tag   string
	Force string
}

type GetAccountByPUUIDParams struct {
	Puuid string
	Force string
}

type GetAccountResponse struct {
	Status int `json:"status"`
	Data   struct {
		Puuid        string `json:"puuid"`
		Region       string `json:"region"`
		AccountLevel int    `json:"account_level"`
		Name         string `json:"name"`
		Tag          string `json:"tag"`
		Card         struct {
			Small string `json:"small"`
			Large string `json:"large"`
			Wide  string `json:"wide"`
			ID    string `json:"id"`
		} `json:"card"`
		LastUpdate    string `json:"last_update"`
		LastUpdateRaw int    `json:"last_update_raw"`
	} `json:"data"`
	Errors []Error `json:"errors"`
}

type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Details string `json:"details"`
}

type GetLifetimeMatchesByPUUIDParams struct {
	Affinity string
	PUUID    string
	Mode     string
	Map      string
	Page     string
	Size     string
}

type GetLifetimeMatchesByPUUIDResponse struct {
	Status  int    `json:"status"`
	Name    string `json:"name"`
	Tag     string `json:"tag"`
	Results struct {
		Total    int `json:"total"`
		Returned int `json:"returned"`
		Before   int `json:"before"`
		After    int `json:"after"`
	} `json:"results"`
	Data []struct {
		Meta struct {
			ID  string `json:"id"`
			Map struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"map"`
			Version   string `json:"version"`
			Mode      string `json:"mode"`
			StartedAt string `json:"started_at"`
			Season    struct {
				ID    string `json:"id"`
				Short string `json:"short"`
			} `json:"season"`
			Region  string `json:"region"`
			Cluster string `json:"cluster"`
		} `json:"meta"`
		Stats struct {
			Puuid     string `json:"puuid"`
			Team      string `json:"team"`
			Level     int    `json:"level"`
			Character struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"character"`
			Tier    int `json:"tier"`
			Score   int `json:"score"`
			Kills   int `json:"kills"`
			Deaths  int `json:"deaths"`
			Assists int `json:"assists"`
			Shots   struct {
				Head int `json:"head"`
				Body int `json:"body"`
				Leg  int `json:"leg"`
			} `json:"shots"`
			Damage struct {
				Made     int `json:"made"`
				Received int `json:"received"`
			} `json:"damage"`
		} `json:"stats"`
		Teams struct {
			Red  int `json:"red"`
			Blue int `json:"blue"`
		} `json:"teams"`
	} `json:"data"`
	Errors []Error `json:"errors"`
}

type GetLifetimeMMRHistoryByPUUIDParams struct {
	Affinity string
	Puuid    string
	Page     string
	Size     string
}

type GetLifetimeMMRHistoryByPUUIDResponse struct {
	Status  int    `json:"status"`
	Name    string `json:"name"`
	Tag     string `json:"tag"`
	Results struct {
		Total    int `json:"total"`
		Returned int `json:"returned"`
		Before   int `json:"before"`
		After    int `json:"after"`
	} `json:"results"`
	Data []struct {
		MatchID string `json:"match_id"`
		Tier    struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"tier"`
		Map struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"map"`
		Season struct {
			ID    string `json:"id"`
			Short string `json:"short"`
		} `json:"season"`
		RankingInTier int       `json:"ranking_in_tier"`
		LastMmrChange int       `json:"last_mmr_change"`
		Elo           int       `json:"elo"`
		Date          time.Time `json:"date"`
	} `json:"data"`
	Errors []Error `json:"errors"`
}

type GetMatchesByPUUIDv3Params struct {
	Affinity string
	Puuid    string
	Mode     string
	Map      string
	Size     string
}

type GetMatchesByPUUIDv3Response struct {
	Status int `json:"status"`
	Data   []struct {
		Metadata struct {
			Map              string `json:"map"`
			GameVersion      string `json:"game_version"`
			GameLength       int    `json:"game_length"`
			GameStart        int    `json:"game_start"`
			GameStartPatched string `json:"game_start_patched"`
			RoundsPlayed     int    `json:"rounds_played"`
			Mode             string `json:"mode"`
			ModeID           string `json:"mode_id"`
			Queue            string `json:"queue"`
			SeasonID         string `json:"season_id"`
			Platform         string `json:"platform"`
			Matchid          string `json:"matchid"`
			PremierInfo      struct {
				TournamentID string `json:"tournament_id"`
				MatchupID    string `json:"matchup_id"`
			} `json:"premier_info"`
			Region  string `json:"region"`
			Cluster string `json:"cluster"`
		} `json:"metadata"`
		Players struct {
			AllPlayers []struct {
				Puuid              string `json:"puuid"`
				Name               string `json:"name"`
				Tag                string `json:"tag"`
				Team               string `json:"team"`
				Level              int    `json:"level"`
				Character          string `json:"character"`
				Currenttier        int    `json:"currenttier"`
				CurrenttierPatched string `json:"currenttier_patched"`
				PlayerCard         string `json:"player_card"`
				PlayerTitle        string `json:"player_title"`
				PartyID            string `json:"party_id"`
				SessionPlaytime    struct {
					Minutes      int `json:"minutes"`
					Seconds      int `json:"seconds"`
					Milliseconds int `json:"milliseconds"`
				} `json:"session_playtime"`
				Assets struct {
					Card struct {
						Small string `json:"small"`
						Large string `json:"large"`
						Wide  string `json:"wide"`
					} `json:"card"`
					Agent struct {
						Small    string `json:"small"`
						Full     string `json:"full"`
						Bust     string `json:"bust"`
						Killfeed string `json:"killfeed"`
					} `json:"agent"`
				} `json:"assets"`
				Behaviour struct {
					AfkRounds    int `json:"afk_rounds"`
					FriendlyFire struct {
						Incoming int `json:"incoming"`
						Outgoing int `json:"outgoing"`
					} `json:"friendly_fire"`
					RoundsInSpawn int `json:"rounds_in_spawn"`
				} `json:"behaviour"`
				Platform struct {
					Type string `json:"type"`
					Os   struct {
						Name    string `json:"name"`
						Version string `json:"version"`
					} `json:"os"`
				} `json:"platform"`
				AbilityCasts struct {
					CCast int `json:"c_cast"`
					QCast int `json:"q_cast"`
					ECast int `json:"e_cast"`
					XCast int `json:"x_cast"`
				} `json:"ability_casts"`
				Stats struct {
					Score     int `json:"score"`
					Kills     int `json:"kills"`
					Deaths    int `json:"deaths"`
					Assists   int `json:"assists"`
					Bodyshots int `json:"bodyshots"`
					Headshots int `json:"headshots"`
					Legshots  int `json:"legshots"`
				} `json:"stats"`
				Economy struct {
					Spent struct {
						Overall int `json:"overall"`
						Average int `json:"average"`
					} `json:"spent"`
					LoadoutValue struct {
						Overall int `json:"overall"`
						Average int `json:"average"`
					} `json:"loadout_value"`
				} `json:"economy"`
				DamageMade     int `json:"damage_made"`
				DamageReceived int `json:"damage_received"`
			} `json:"all_players"`
			Red []struct {
				Puuid              string `json:"puuid"`
				Name               string `json:"name"`
				Tag                string `json:"tag"`
				Team               string `json:"team"`
				Level              int    `json:"level"`
				Character          string `json:"character"`
				Currenttier        int    `json:"currenttier"`
				CurrenttierPatched string `json:"currenttier_patched"`
				PlayerCard         string `json:"player_card"`
				PlayerTitle        string `json:"player_title"`
				PartyID            string `json:"party_id"`
				SessionPlaytime    struct {
					Minutes      int `json:"minutes"`
					Seconds      int `json:"seconds"`
					Milliseconds int `json:"milliseconds"`
				} `json:"session_playtime"`
				Assets struct {
					Card struct {
						Small string `json:"small"`
						Large string `json:"large"`
						Wide  string `json:"wide"`
					} `json:"card"`
					Agent struct {
						Small    string `json:"small"`
						Full     string `json:"full"`
						Bust     string `json:"bust"`
						Killfeed string `json:"killfeed"`
					} `json:"agent"`
				} `json:"assets"`
				Behaviour struct {
					AfkRounds    int `json:"afk_rounds"`
					FriendlyFire struct {
						Incoming int `json:"incoming"`
						Outgoing int `json:"outgoing"`
					} `json:"friendly_fire"`
					RoundsInSpawn int `json:"rounds_in_spawn"`
				} `json:"behaviour"`
				Platform struct {
					Type string `json:"type"`
					Os   struct {
						Name    string `json:"name"`
						Version string `json:"version"`
					} `json:"os"`
				} `json:"platform"`
				AbilityCasts struct {
					CCast int `json:"c_cast"`
					QCast int `json:"q_cast"`
					ECast int `json:"e_cast"`
					XCast int `json:"x_cast"`
				} `json:"ability_casts"`
				Stats struct {
					Score     int `json:"score"`
					Kills     int `json:"kills"`
					Deaths    int `json:"deaths"`
					Assists   int `json:"assists"`
					Bodyshots int `json:"bodyshots"`
					Headshots int `json:"headshots"`
					Legshots  int `json:"legshots"`
				} `json:"stats"`
				Economy struct {
					Spent struct {
						Overall int `json:"overall"`
						Average int `json:"average"`
					} `json:"spent"`
					LoadoutValue struct {
						Overall int `json:"overall"`
						Average int `json:"average"`
					} `json:"loadout_value"`
				} `json:"economy"`
				DamageMade     int `json:"damage_made"`
				DamageReceived int `json:"damage_received"`
			} `json:"red"`
			Blue []struct {
				Puuid              string `json:"puuid"`
				Name               string `json:"name"`
				Tag                string `json:"tag"`
				Team               string `json:"team"`
				Level              int    `json:"level"`
				Character          string `json:"character"`
				Currenttier        int    `json:"currenttier"`
				CurrenttierPatched string `json:"currenttier_patched"`
				PlayerCard         string `json:"player_card"`
				PlayerTitle        string `json:"player_title"`
				PartyID            string `json:"party_id"`
				SessionPlaytime    struct {
					Minutes      int `json:"minutes"`
					Seconds      int `json:"seconds"`
					Milliseconds int `json:"milliseconds"`
				} `json:"session_playtime"`
				Assets struct {
					Card struct {
						Small string `json:"small"`
						Large string `json:"large"`
						Wide  string `json:"wide"`
					} `json:"card"`
					Agent struct {
						Small    string `json:"small"`
						Full     string `json:"full"`
						Bust     string `json:"bust"`
						Killfeed string `json:"killfeed"`
					} `json:"agent"`
				} `json:"assets"`
				Behaviour struct {
					AfkRounds    int `json:"afk_rounds"`
					FriendlyFire struct {
						Incoming int `json:"incoming"`
						Outgoing int `json:"outgoing"`
					} `json:"friendly_fire"`
					RoundsInSpawn int `json:"rounds_in_spawn"`
				} `json:"behaviour"`
				Platform struct {
					Type string `json:"type"`
					Os   struct {
						Name    string `json:"name"`
						Version string `json:"version"`
					} `json:"os"`
				} `json:"platform"`
				AbilityCasts struct {
					CCast int `json:"c_cast"`
					QCast int `json:"q_cast"`
					ECast int `json:"e_cast"`
					XCast int `json:"x_cast"`
				} `json:"ability_casts"`
				Stats struct {
					Score     int `json:"score"`
					Kills     int `json:"kills"`
					Deaths    int `json:"deaths"`
					Assists   int `json:"assists"`
					Bodyshots int `json:"bodyshots"`
					Headshots int `json:"headshots"`
					Legshots  int `json:"legshots"`
				} `json:"stats"`
				Economy struct {
					Spent struct {
						Overall int `json:"overall"`
						Average int `json:"average"`
					} `json:"spent"`
					LoadoutValue struct {
						Overall int `json:"overall"`
						Average int `json:"average"`
					} `json:"loadout_value"`
				} `json:"economy"`
				DamageMade     int `json:"damage_made"`
				DamageReceived int `json:"damage_received"`
			} `json:"blue"`
		} `json:"players"`
		Observers []struct {
			Puuid    string `json:"puuid"`
			Name     string `json:"name"`
			Tag      string `json:"tag"`
			Platform struct {
				Type string `json:"type"`
				Os   struct {
					Name    string `json:"name"`
					Version string `json:"version"`
				} `json:"os"`
			} `json:"platform"`
			SessionPlaytime struct {
				Minutes      int `json:"minutes"`
				Seconds      int `json:"seconds"`
				Milliseconds int `json:"milliseconds"`
			} `json:"session_playtime"`
			Team        string `json:"team"`
			Level       int    `json:"level"`
			PlayerCard  string `json:"player_card"`
			PlayerTitle string `json:"player_title"`
			PartyID     string `json:"party_id"`
		} `json:"observers"`
		Coaches []struct {
			Puuid string `json:"puuid"`
			Team  string `json:"team"`
		} `json:"coaches"`
		Teams struct {
			Red struct {
				HasWon     bool `json:"has_won"`
				RoundsWon  int  `json:"rounds_won"`
				RoundsLost int  `json:"rounds_lost"`
				Roaster    struct {
					Members       []string `json:"members"`
					Name          string   `json:"name"`
					Tag           string   `json:"tag"`
					Customization struct {
						Icon      string `json:"icon"`
						Image     string `json:"image"`
						Primary   string `json:"primary"`
						Secondary string `json:"secondary"`
						Tertiary  string `json:"tertiary"`
					} `json:"customization"`
				} `json:"roaster"`
			} `json:"red"`
			Blue struct {
				HasWon     bool `json:"has_won"`
				RoundsWon  int  `json:"rounds_won"`
				RoundsLost int  `json:"rounds_lost"`
				Roaster    struct {
					Members       []string `json:"members"`
					Name          string   `json:"name"`
					Tag           string   `json:"tag"`
					Customization struct {
						Icon      string `json:"icon"`
						Image     string `json:"image"`
						Primary   string `json:"primary"`
						Secondary string `json:"secondary"`
						Tertiary  string `json:"tertiary"`
					} `json:"customization"`
				} `json:"roaster"`
			} `json:"blue"`
		} `json:"teams"`
		Rounds []struct {
			WinningTeam string `json:"winning_team"`
			EndType     string `json:"end_type"`
			BombPlanted bool   `json:"bomb_planted"`
			BombDefused bool   `json:"bomb_defused"`
			PlantEvents struct {
				PlantLocation struct {
					X int `json:"x"`
					Y int `json:"y"`
				} `json:"plant_location"`
				PlantedBy struct {
					Puuid       string `json:"puuid"`
					DisplayName string `json:"display_name"`
					Team        string `json:"team"`
				} `json:"planted_by"`
				PlantSite              string `json:"plant_site"`
				PlantTimeInRound       int    `json:"plant_time_in_round"`
				PlayerLocationsOnPlant []struct {
					PlayerPuuid       string `json:"player_puuid"`
					PlayerDisplayName string `json:"player_display_name"`
					PlayerTeam        string `json:"player_team"`
					Location          struct {
						X int `json:"x"`
						Y int `json:"y"`
					} `json:"location"`
					ViewRadians float64 `json:"view_radians"`
				} `json:"player_locations_on_plant"`
			} `json:"plant_events"`
			DefuseEvents struct {
				DefuseLocation struct {
					X int `json:"x"`
					Y int `json:"y"`
				} `json:"defuse_location"`
				DefusedBy struct {
					Puuid       string `json:"puuid"`
					DisplayName string `json:"display_name"`
					Team        string `json:"team"`
				} `json:"defused_by"`
				DefuseTimeInRound       int `json:"defuse_time_in_round"`
				PlayerLocationsOnDefuse []struct {
					PlayerPuuid       string `json:"player_puuid"`
					PlayerDisplayName string `json:"player_display_name"`
					PlayerTeam        string `json:"player_team"`
					Location          struct {
						X int `json:"x"`
						Y int `json:"y"`
					} `json:"location"`
					ViewRadians float64 `json:"view_radians"`
				} `json:"player_locations_on_defuse"`
			} `json:"defuse_events"`
			PlayerStats []struct {
				AbilityCasts struct {
					CCasts int `json:"c_casts"`
					QCasts int `json:"q_casts"`
					ECasts int `json:"e_casts"`
					XCasts int `json:"x_casts"`
				} `json:"ability_casts"`
				PlayerPuuid       string        `json:"player_puuid"`
				PlayerDisplayName string        `json:"player_display_name"`
				PlayerTeam        string        `json:"player_team"`
				DamageEvents      []interface{} `json:"damage_events"`
				Damage            int           `json:"damage"`
				Bodyshots         int           `json:"bodyshots"`
				Headshots         int           `json:"headshots"`
				Legshots          int           `json:"legshots"`
				KillEvents        []interface{} `json:"kill_events"`
				Kills             int           `json:"kills"`
				Score             int           `json:"score"`
				Economy           struct {
					LoadoutValue int `json:"loadout_value"`
					Weapon       struct {
						ID     string `json:"id"`
						Name   string `json:"name"`
						Assets struct {
							DisplayIcon  string `json:"display_icon"`
							KillfeedIcon string `json:"killfeed_icon"`
						} `json:"assets"`
					} `json:"weapon"`
					Armor struct {
						ID     string `json:"id"`
						Name   string `json:"name"`
						Assets struct {
							DisplayIcon string `json:"display_icon"`
						} `json:"assets"`
					} `json:"armor"`
					Remaining int `json:"remaining"`
					Spent     int `json:"spent"`
				} `json:"economy"`
				WasAfk        bool `json:"was_afk"`
				WasPenalized  bool `json:"was_penalized"`
				StayedInSpawn bool `json:"stayed_in_spawn"`
			} `json:"player_stats"`
		} `json:"rounds"`
	} `json:"data"`
	Errors []Error `json:"errors"`
}

type GetMMRByPUUIDv2Params struct {
	Affinity string
	Puuid    string
	Season   string
}

type GetMMRByPUUIDv2Response struct {
	Status int `json:"status"`
	Data   struct {
		Name        string `json:"name"`
		Tag         string `json:"tag"`
		CurrentData struct {
			Currenttier        int    `json:"currenttier"`
			CurrenttierPatched string `json:"currenttier_patched"`
			Images             struct {
				Small        string `json:"small"`
				Large        string `json:"large"`
				TriangleDown string `json:"triangle_down"`
				TriangleUp   string `json:"triangle_up"`
			} `json:"images"`
			RankingInTier       int  `json:"ranking_in_tier"`
			MmrChangeToLastGame int  `json:"mmr_change_to_last_game"`
			Elo                 int  `json:"elo"`
			Old                 bool `json:"old"`
		} `json:"current_data"`
		HighestRank struct {
			Old         bool   `json:"old"`
			Tier        int    `json:"tier"`
			PatchedTier string `json:"patched_tier"`
			Season      string `json:"season"`
		} `json:"highest_rank"`
		BySeason struct {
			E6A3 struct {
				Error            string `json:"error"`
				Wins             int    `json:"wins"`
				NumberOfGames    int    `json:"number_of_games"`
				FinalRank        int    `json:"final_rank"`
				FinalRankPatched string `json:"final_rank_patched"`
				ActRankWins      []struct {
					PatchedTier string `json:"patched_tier"`
					Tier        int    `json:"tier"`
				} `json:"act_rank_wins"`
				Old bool `json:"old"`
			} `json:"e6a3"`
			E6A2 struct {
				Error            string `json:"error"`
				Wins             int    `json:"wins"`
				NumberOfGames    int    `json:"number_of_games"`
				FinalRank        int    `json:"final_rank"`
				FinalRankPatched string `json:"final_rank_patched"`
				ActRankWins      []struct {
					PatchedTier string `json:"patched_tier"`
					Tier        int    `json:"tier"`
				} `json:"act_rank_wins"`
				Old bool `json:"old"`
			} `json:"e6a2"`
			E6A1 struct {
				Error            string `json:"error"`
				Wins             int    `json:"wins"`
				NumberOfGames    int    `json:"number_of_games"`
				FinalRank        int    `json:"final_rank"`
				FinalRankPatched string `json:"final_rank_patched"`
				ActRankWins      []struct {
					PatchedTier string `json:"patched_tier"`
					Tier        int    `json:"tier"`
				} `json:"act_rank_wins"`
				Old bool `json:"old"`
			} `json:"e6a1"`
			E5A3 struct {
				Error            string `json:"error"`
				Wins             int    `json:"wins"`
				NumberOfGames    int    `json:"number_of_games"`
				FinalRank        int    `json:"final_rank"`
				FinalRankPatched string `json:"final_rank_patched"`
				ActRankWins      []struct {
					PatchedTier string `json:"patched_tier"`
					Tier        int    `json:"tier"`
				} `json:"act_rank_wins"`
				Old bool `json:"old"`
			} `json:"e5a3"`
			E5A2 struct {
				Error            string `json:"error"`
				Wins             int    `json:"wins"`
				NumberOfGames    int    `json:"number_of_games"`
				FinalRank        int    `json:"final_rank"`
				FinalRankPatched string `json:"final_rank_patched"`
				ActRankWins      []struct {
					PatchedTier string `json:"patched_tier"`
					Tier        int    `json:"tier"`
				} `json:"act_rank_wins"`
				Old bool `json:"old"`
			} `json:"e5a2"`
			E5A1 struct {
				Error            string `json:"error"`
				Wins             int    `json:"wins"`
				NumberOfGames    int    `json:"number_of_games"`
				FinalRank        int    `json:"final_rank"`
				FinalRankPatched string `json:"final_rank_patched"`
				ActRankWins      []struct {
					PatchedTier string `json:"patched_tier"`
					Tier        int    `json:"tier"`
				} `json:"act_rank_wins"`
				Old bool `json:"old"`
			} `json:"e5a1"`
			E4A3 struct {
				Error            string `json:"error"`
				Wins             int    `json:"wins"`
				NumberOfGames    int    `json:"number_of_games"`
				FinalRank        int    `json:"final_rank"`
				FinalRankPatched string `json:"final_rank_patched"`
				ActRankWins      []struct {
					PatchedTier string `json:"patched_tier"`
					Tier        int    `json:"tier"`
				} `json:"act_rank_wins"`
				Old bool `json:"old"`
			} `json:"e4a3"`
			E4A2 struct {
				Error            string `json:"error"`
				Wins             int    `json:"wins"`
				NumberOfGames    int    `json:"number_of_games"`
				FinalRank        int    `json:"final_rank"`
				FinalRankPatched string `json:"final_rank_patched"`
				ActRankWins      []struct {
					PatchedTier string `json:"patched_tier"`
					Tier        int    `json:"tier"`
				} `json:"act_rank_wins"`
				Old bool `json:"old"`
			} `json:"e4a2"`
			E4A1 struct {
				Error            string `json:"error"`
				Wins             int    `json:"wins"`
				NumberOfGames    int    `json:"number_of_games"`
				FinalRank        int    `json:"final_rank"`
				FinalRankPatched string `json:"final_rank_patched"`
				ActRankWins      []struct {
					PatchedTier string `json:"patched_tier"`
					Tier        int    `json:"tier"`
				} `json:"act_rank_wins"`
				Old bool `json:"old"`
			} `json:"e4a1"`
			E3A3 struct {
				Error            string `json:"error"`
				Wins             int    `json:"wins"`
				NumberOfGames    int    `json:"number_of_games"`
				FinalRank        int    `json:"final_rank"`
				FinalRankPatched string `json:"final_rank_patched"`
				ActRankWins      []struct {
					PatchedTier string `json:"patched_tier"`
					Tier        int    `json:"tier"`
				} `json:"act_rank_wins"`
				Old bool `json:"old"`
			} `json:"e3a3"`
			E3A2 struct {
				Error            string `json:"error"`
				Wins             int    `json:"wins"`
				NumberOfGames    int    `json:"number_of_games"`
				FinalRank        int    `json:"final_rank"`
				FinalRankPatched string `json:"final_rank_patched"`
				ActRankWins      []struct {
					PatchedTier string `json:"patched_tier"`
					Tier        int    `json:"tier"`
				} `json:"act_rank_wins"`
				Old bool `json:"old"`
			} `json:"e3a2"`
			E3A1 struct {
				Error            string `json:"error"`
				Wins             int    `json:"wins"`
				NumberOfGames    int    `json:"number_of_games"`
				FinalRank        int    `json:"final_rank"`
				FinalRankPatched string `json:"final_rank_patched"`
				ActRankWins      []struct {
					PatchedTier string `json:"patched_tier"`
					Tier        int    `json:"tier"`
				} `json:"act_rank_wins"`
				Old bool `json:"old"`
			} `json:"e3a1"`
			E2A3 struct {
				Error            string `json:"error"`
				Wins             int    `json:"wins"`
				NumberOfGames    int    `json:"number_of_games"`
				FinalRank        int    `json:"final_rank"`
				FinalRankPatched string `json:"final_rank_patched"`
				ActRankWins      []struct {
					PatchedTier string `json:"patched_tier"`
					Tier        int    `json:"tier"`
				} `json:"act_rank_wins"`
				Old bool `json:"old"`
			} `json:"e2a3"`
			E2A2 struct {
				Error            string `json:"error"`
				Wins             int    `json:"wins"`
				NumberOfGames    int    `json:"number_of_games"`
				FinalRank        int    `json:"final_rank"`
				FinalRankPatched string `json:"final_rank_patched"`
				ActRankWins      []struct {
					PatchedTier string `json:"patched_tier"`
					Tier        int    `json:"tier"`
				} `json:"act_rank_wins"`
				Old bool `json:"old"`
			} `json:"e2a2"`
			E2A1 struct {
				Error            string `json:"error"`
				Wins             int    `json:"wins"`
				NumberOfGames    int    `json:"number_of_games"`
				FinalRank        int    `json:"final_rank"`
				FinalRankPatched string `json:"final_rank_patched"`
				ActRankWins      []struct {
					PatchedTier string `json:"patched_tier"`
					Tier        int    `json:"tier"`
				} `json:"act_rank_wins"`
				Old bool `json:"old"`
			} `json:"e2a1"`
			E1A3 struct {
				Error            string `json:"error"`
				Wins             int    `json:"wins"`
				NumberOfGames    int    `json:"number_of_games"`
				FinalRank        int    `json:"final_rank"`
				FinalRankPatched string `json:"final_rank_patched"`
				ActRankWins      []struct {
					PatchedTier string `json:"patched_tier"`
					Tier        int    `json:"tier"`
				} `json:"act_rank_wins"`
				Old bool `json:"old"`
			} `json:"e1a3"`
			E1A2 struct {
				Error            string `json:"error"`
				Wins             int    `json:"wins"`
				NumberOfGames    int    `json:"number_of_games"`
				FinalRank        int    `json:"final_rank"`
				FinalRankPatched string `json:"final_rank_patched"`
				ActRankWins      []struct {
					PatchedTier string `json:"patched_tier"`
					Tier        int    `json:"tier"`
				} `json:"act_rank_wins"`
				Old bool `json:"old"`
			} `json:"e1a2"`
			E1A1 struct {
				Error            string `json:"error"`
				Wins             int    `json:"wins"`
				NumberOfGames    int    `json:"number_of_games"`
				FinalRank        int    `json:"final_rank"`
				FinalRankPatched string `json:"final_rank_patched"`
				ActRankWins      []struct {
					PatchedTier string `json:"patched_tier"`
					Tier        int    `json:"tier"`
				} `json:"act_rank_wins"`
				Old bool `json:"old"`
			} `json:"e1a1"`
		} `json:"by_season"`
		Wins             int    `json:"wins"`
		NumberOfGames    int    `json:"number_of_games"`
		FinalRank        int    `json:"final_rank"`
		FinalRankPatched string `json:"final_rank_patched"`
		ActRankWins      []struct {
			PatchedTier string `json:"patched_tier"`
			Tier        int    `json:"tier"`
		} `json:"act_rank_wins"`
		Old bool `json:"old"`
	} `json:"data"`
	Errors []Error `json:"errors"`
}

type GetMMRHistoryByPUUIDParams struct {
	Affinity string
	Puuid    string
}

type GetMMRHistoryByPUUIDResponse struct {
	Status int    `json:"status"`
	Name   string `json:"name"`
	Tag    string `json:"tag"`
	Data   []struct {
		Currenttier        int    `json:"currenttier"`
		CurrenttierPatched string `json:"currenttier_patched"`
		Images             struct {
			Small        string `json:"small"`
			Large        string `json:"large"`
			TriangleDown string `json:"triangle_down"`
			TriangleUp   string `json:"triangle_up"`
		} `json:"images"`
		MatchID string `json:"match_id"`
		Map     struct {
			Name string `json:"name"`
			ID   string `json:"id"`
		} `json:"map"`
		SeasonID            string `json:"season_id"`
		RankingInTier       int    `json:"ranking_in_tier"`
		MmrChangeToLastGame int    `json:"mmr_change_to_last_game"`
		Elo                 int    `json:"elo"`
		Date                string `json:"date"`
		DateRaw             int    `json:"date_raw"`
	} `json:"data"`
	Errors []Error `json:"errors"`
}

type GetContentParams struct {
	Locale string
}

type GetContentResponse struct {
	Version    string `json:"version"`
	Characters []struct {
		Name           string `json:"name"`
		LocalizedNames []struct {
			ArAE string `json:"ar-AE"`
			DeDE string `json:"de-DE"`
			EnGB string `json:"en-GB"`
			EnUS string `json:"en-US"`
			EsES string `json:"es-ES"`
			EsMX string `json:"es-MX"`
			FrFR string `json:"fr-FR"`
			IDID string `json:"id-ID"`
			ItIT string `json:"it-IT"`
			JaJP string `json:"ja-JP"`
			KoKR string `json:"ko-KR"`
			PlPL string `json:"pl-PL"`
			PtBR string `json:"pt-BR"`
			RuRU string `json:"ru-RU"`
			ThTH string `json:"th-TH"`
			TrTR string `json:"tr-TR"`
			ViVN string `json:"vi-VN"`
			ZnCN string `json:"zn-CN"`
			ZnTW string `json:"zn-TW"`
		} `json:"localizedNames"`
		ID        string `json:"id"`
		AssetName string `json:"assetName"`
		AssetPath string `json:"assetPath"`
	} `json:"characters"`
	Maps []struct {
		Name           string `json:"name"`
		LocalizedNames []struct {
			ArAE string `json:"ar-AE"`
			DeDE string `json:"de-DE"`
			EnGB string `json:"en-GB"`
			EnUS string `json:"en-US"`
			EsES string `json:"es-ES"`
			EsMX string `json:"es-MX"`
			FrFR string `json:"fr-FR"`
			IDID string `json:"id-ID"`
			ItIT string `json:"it-IT"`
			JaJP string `json:"ja-JP"`
			KoKR string `json:"ko-KR"`
			PlPL string `json:"pl-PL"`
			PtBR string `json:"pt-BR"`
			RuRU string `json:"ru-RU"`
			ThTH string `json:"th-TH"`
			TrTR string `json:"tr-TR"`
			ViVN string `json:"vi-VN"`
			ZnCN string `json:"zn-CN"`
			ZnTW string `json:"zn-TW"`
		} `json:"localizedNames"`
		ID        string `json:"id"`
		AssetName string `json:"assetName"`
		AssetPath string `json:"assetPath"`
	} `json:"maps"`
	Chromas []struct {
		Name           string `json:"name"`
		LocalizedNames []struct {
			ArAE string `json:"ar-AE"`
			DeDE string `json:"de-DE"`
			EnGB string `json:"en-GB"`
			EnUS string `json:"en-US"`
			EsES string `json:"es-ES"`
			EsMX string `json:"es-MX"`
			FrFR string `json:"fr-FR"`
			IDID string `json:"id-ID"`
			ItIT string `json:"it-IT"`
			JaJP string `json:"ja-JP"`
			KoKR string `json:"ko-KR"`
			PlPL string `json:"pl-PL"`
			PtBR string `json:"pt-BR"`
			RuRU string `json:"ru-RU"`
			ThTH string `json:"th-TH"`
			TrTR string `json:"tr-TR"`
			ViVN string `json:"vi-VN"`
			ZnCN string `json:"zn-CN"`
			ZnTW string `json:"zn-TW"`
		} `json:"localizedNames"`
		ID        string `json:"id"`
		AssetName string `json:"assetName"`
		AssetPath string `json:"assetPath"`
	} `json:"chromas"`
	Skins []struct {
		Name           string `json:"name"`
		LocalizedNames []struct {
			ArAE string `json:"ar-AE"`
			DeDE string `json:"de-DE"`
			EnGB string `json:"en-GB"`
			EnUS string `json:"en-US"`
			EsES string `json:"es-ES"`
			EsMX string `json:"es-MX"`
			FrFR string `json:"fr-FR"`
			IDID string `json:"id-ID"`
			ItIT string `json:"it-IT"`
			JaJP string `json:"ja-JP"`
			KoKR string `json:"ko-KR"`
			PlPL string `json:"pl-PL"`
			PtBR string `json:"pt-BR"`
			RuRU string `json:"ru-RU"`
			ThTH string `json:"th-TH"`
			TrTR string `json:"tr-TR"`
			ViVN string `json:"vi-VN"`
			ZnCN string `json:"zn-CN"`
			ZnTW string `json:"zn-TW"`
		} `json:"localizedNames"`
		ID        string `json:"id"`
		AssetName string `json:"assetName"`
		AssetPath string `json:"assetPath"`
	} `json:"skins"`
	SkinLevels []struct {
		Name           string `json:"name"`
		LocalizedNames []struct {
			ArAE string `json:"ar-AE"`
			DeDE string `json:"de-DE"`
			EnGB string `json:"en-GB"`
			EnUS string `json:"en-US"`
			EsES string `json:"es-ES"`
			EsMX string `json:"es-MX"`
			FrFR string `json:"fr-FR"`
			IDID string `json:"id-ID"`
			ItIT string `json:"it-IT"`
			JaJP string `json:"ja-JP"`
			KoKR string `json:"ko-KR"`
			PlPL string `json:"pl-PL"`
			PtBR string `json:"pt-BR"`
			RuRU string `json:"ru-RU"`
			ThTH string `json:"th-TH"`
			TrTR string `json:"tr-TR"`
			ViVN string `json:"vi-VN"`
			ZnCN string `json:"zn-CN"`
			ZnTW string `json:"zn-TW"`
		} `json:"localizedNames"`
		ID        string `json:"id"`
		AssetName string `json:"assetName"`
		AssetPath string `json:"assetPath"`
	} `json:"skinLevels"`
	Equips []struct {
		Name           string `json:"name"`
		LocalizedNames []struct {
			ArAE string `json:"ar-AE"`
			DeDE string `json:"de-DE"`
			EnGB string `json:"en-GB"`
			EnUS string `json:"en-US"`
			EsES string `json:"es-ES"`
			EsMX string `json:"es-MX"`
			FrFR string `json:"fr-FR"`
			IDID string `json:"id-ID"`
			ItIT string `json:"it-IT"`
			JaJP string `json:"ja-JP"`
			KoKR string `json:"ko-KR"`
			PlPL string `json:"pl-PL"`
			PtBR string `json:"pt-BR"`
			RuRU string `json:"ru-RU"`
			ThTH string `json:"th-TH"`
			TrTR string `json:"tr-TR"`
			ViVN string `json:"vi-VN"`
			ZnCN string `json:"zn-CN"`
			ZnTW string `json:"zn-TW"`
		} `json:"localizedNames"`
		ID        string `json:"id"`
		AssetName string `json:"assetName"`
		AssetPath string `json:"assetPath"`
	} `json:"equips"`
	GameModes []struct {
		Name           string `json:"name"`
		LocalizedNames []struct {
			ArAE string `json:"ar-AE"`
			DeDE string `json:"de-DE"`
			EnGB string `json:"en-GB"`
			EnUS string `json:"en-US"`
			EsES string `json:"es-ES"`
			EsMX string `json:"es-MX"`
			FrFR string `json:"fr-FR"`
			IDID string `json:"id-ID"`
			ItIT string `json:"it-IT"`
			JaJP string `json:"ja-JP"`
			KoKR string `json:"ko-KR"`
			PlPL string `json:"pl-PL"`
			PtBR string `json:"pt-BR"`
			RuRU string `json:"ru-RU"`
			ThTH string `json:"th-TH"`
			TrTR string `json:"tr-TR"`
			ViVN string `json:"vi-VN"`
			ZnCN string `json:"zn-CN"`
			ZnTW string `json:"zn-TW"`
		} `json:"localizedNames"`
		ID        string `json:"id"`
		AssetName string `json:"assetName"`
		AssetPath string `json:"assetPath"`
	} `json:"gameModes"`
	Sprays []struct {
		Name           string `json:"name"`
		LocalizedNames []struct {
			ArAE string `json:"ar-AE"`
			DeDE string `json:"de-DE"`
			EnGB string `json:"en-GB"`
			EnUS string `json:"en-US"`
			EsES string `json:"es-ES"`
			EsMX string `json:"es-MX"`
			FrFR string `json:"fr-FR"`
			IDID string `json:"id-ID"`
			ItIT string `json:"it-IT"`
			JaJP string `json:"ja-JP"`
			KoKR string `json:"ko-KR"`
			PlPL string `json:"pl-PL"`
			PtBR string `json:"pt-BR"`
			RuRU string `json:"ru-RU"`
			ThTH string `json:"th-TH"`
			TrTR string `json:"tr-TR"`
			ViVN string `json:"vi-VN"`
			ZnCN string `json:"zn-CN"`
			ZnTW string `json:"zn-TW"`
		} `json:"localizedNames"`
		ID        string `json:"id"`
		AssetName string `json:"assetName"`
		AssetPath string `json:"assetPath"`
	} `json:"sprays"`
	SprayLevels []struct {
		Name           string `json:"name"`
		LocalizedNames []struct {
			ArAE string `json:"ar-AE"`
			DeDE string `json:"de-DE"`
			EnGB string `json:"en-GB"`
			EnUS string `json:"en-US"`
			EsES string `json:"es-ES"`
			EsMX string `json:"es-MX"`
			FrFR string `json:"fr-FR"`
			IDID string `json:"id-ID"`
			ItIT string `json:"it-IT"`
			JaJP string `json:"ja-JP"`
			KoKR string `json:"ko-KR"`
			PlPL string `json:"pl-PL"`
			PtBR string `json:"pt-BR"`
			RuRU string `json:"ru-RU"`
			ThTH string `json:"th-TH"`
			TrTR string `json:"tr-TR"`
			ViVN string `json:"vi-VN"`
			ZnCN string `json:"zn-CN"`
			ZnTW string `json:"zn-TW"`
		} `json:"localizedNames"`
		ID        string `json:"id"`
		AssetName string `json:"assetName"`
		AssetPath string `json:"assetPath"`
	} `json:"sprayLevels"`
	Charms []struct {
		Name           string `json:"name"`
		LocalizedNames []struct {
			ArAE string `json:"ar-AE"`
			DeDE string `json:"de-DE"`
			EnGB string `json:"en-GB"`
			EnUS string `json:"en-US"`
			EsES string `json:"es-ES"`
			EsMX string `json:"es-MX"`
			FrFR string `json:"fr-FR"`
			IDID string `json:"id-ID"`
			ItIT string `json:"it-IT"`
			JaJP string `json:"ja-JP"`
			KoKR string `json:"ko-KR"`
			PlPL string `json:"pl-PL"`
			PtBR string `json:"pt-BR"`
			RuRU string `json:"ru-RU"`
			ThTH string `json:"th-TH"`
			TrTR string `json:"tr-TR"`
			ViVN string `json:"vi-VN"`
			ZnCN string `json:"zn-CN"`
			ZnTW string `json:"zn-TW"`
		} `json:"localizedNames"`
		ID        string `json:"id"`
		AssetName string `json:"assetName"`
		AssetPath string `json:"assetPath"`
	} `json:"charms"`
	CharmLevels []struct {
		Name           string `json:"name"`
		LocalizedNames []struct {
			ArAE string `json:"ar-AE"`
			DeDE string `json:"de-DE"`
			EnGB string `json:"en-GB"`
			EnUS string `json:"en-US"`
			EsES string `json:"es-ES"`
			EsMX string `json:"es-MX"`
			FrFR string `json:"fr-FR"`
			IDID string `json:"id-ID"`
			ItIT string `json:"it-IT"`
			JaJP string `json:"ja-JP"`
			KoKR string `json:"ko-KR"`
			PlPL string `json:"pl-PL"`
			PtBR string `json:"pt-BR"`
			RuRU string `json:"ru-RU"`
			ThTH string `json:"th-TH"`
			TrTR string `json:"tr-TR"`
			ViVN string `json:"vi-VN"`
			ZnCN string `json:"zn-CN"`
			ZnTW string `json:"zn-TW"`
		} `json:"localizedNames"`
		ID        string `json:"id"`
		AssetName string `json:"assetName"`
		AssetPath string `json:"assetPath"`
	} `json:"charmLevels"`
	PlayerCards []struct {
		Name           string `json:"name"`
		LocalizedNames []struct {
			ArAE string `json:"ar-AE"`
			DeDE string `json:"de-DE"`
			EnGB string `json:"en-GB"`
			EnUS string `json:"en-US"`
			EsES string `json:"es-ES"`
			EsMX string `json:"es-MX"`
			FrFR string `json:"fr-FR"`
			IDID string `json:"id-ID"`
			ItIT string `json:"it-IT"`
			JaJP string `json:"ja-JP"`
			KoKR string `json:"ko-KR"`
			PlPL string `json:"pl-PL"`
			PtBR string `json:"pt-BR"`
			RuRU string `json:"ru-RU"`
			ThTH string `json:"th-TH"`
			TrTR string `json:"tr-TR"`
			ViVN string `json:"vi-VN"`
			ZnCN string `json:"zn-CN"`
			ZnTW string `json:"zn-TW"`
		} `json:"localizedNames"`
		ID        string `json:"id"`
		AssetName string `json:"assetName"`
		AssetPath string `json:"assetPath"`
	} `json:"playerCards"`
	PlayerTitles []struct {
		Name           string `json:"name"`
		LocalizedNames []struct {
			ArAE string `json:"ar-AE"`
			DeDE string `json:"de-DE"`
			EnGB string `json:"en-GB"`
			EnUS string `json:"en-US"`
			EsES string `json:"es-ES"`
			EsMX string `json:"es-MX"`
			FrFR string `json:"fr-FR"`
			IDID string `json:"id-ID"`
			ItIT string `json:"it-IT"`
			JaJP string `json:"ja-JP"`
			KoKR string `json:"ko-KR"`
			PlPL string `json:"pl-PL"`
			PtBR string `json:"pt-BR"`
			RuRU string `json:"ru-RU"`
			ThTH string `json:"th-TH"`
			TrTR string `json:"tr-TR"`
			ViVN string `json:"vi-VN"`
			ZnCN string `json:"zn-CN"`
			ZnTW string `json:"zn-TW"`
		} `json:"localizedNames"`
		ID        string `json:"id"`
		AssetName string `json:"assetName"`
		AssetPath string `json:"assetPath"`
	} `json:"playerTitles"`
	Acts []struct {
		Name           string `json:"name"`
		LocalizedNames []struct {
			ArAE string `json:"ar-AE"`
			DeDE string `json:"de-DE"`
			EnGB string `json:"en-GB"`
			EnUS string `json:"en-US"`
			EsES string `json:"es-ES"`
			EsMX string `json:"es-MX"`
			FrFR string `json:"fr-FR"`
			IDID string `json:"id-ID"`
			ItIT string `json:"it-IT"`
			JaJP string `json:"ja-JP"`
			KoKR string `json:"ko-KR"`
			PlPL string `json:"pl-PL"`
			PtBR string `json:"pt-BR"`
			RuRU string `json:"ru-RU"`
			ThTH string `json:"th-TH"`
			TrTR string `json:"tr-TR"`
			ViVN string `json:"vi-VN"`
			ZnCN string `json:"zn-CN"`
			ZnTW string `json:"zn-TW"`
		} `json:"localizedNames"`
		ID       string `json:"id"`
		IsActive bool   `json:"isActive"`
	} `json:"acts"`
	Errors []Error `json:"errors"`
}

type GetEsportsScheduleParams struct {
	Affinity string
	League   string
}

type GetEsportsScheduleResponse struct {
	Status int `json:"status"`
	Data   []struct {
		Date   time.Time `json:"date"`
		State  string    `json:"state"`
		Type   string    `json:"type"`
		Vod    string    `json:"vod"`
		League struct {
			Name       string `json:"name"`
			Identifier string `json:"identifier"`
			Icon       string `json:"icon"`
			Region     string `json:"region"`
		} `json:"league"`
		Tournament struct {
			Name   string `json:"name"`
			Season string `json:"season"`
		} `json:"tournament"`
		Match struct {
			ID       string `json:"id"`
			GameType struct {
				Type  string `json:"type"`
				Count int    `json:"count"`
			} `json:"game_type"`
			Teams []struct {
				Name     string `json:"name"`
				Code     string `json:"code"`
				Icon     string `json:"icon"`
				HasWon   bool   `json:"has_won"`
				GameWins int    `json:"game_wins"`
				Record   struct {
					Wins   int `json:"wins"`
					Losses int `json:"losses"`
				} `json:"record"`
			} `json:"teams"`
		} `json:"match"`
	} `json:"data"`
	Errors []Error `json:"errors"`
}

type GetLeaderboardV2Params struct {
	Affinity string
	Puuid    string
	Name     string
	Tag      string
	Season   string
}

type GetLeaderboardV2Response struct {
	LastUpdate         int `json:"last_update"`
	NextUpdate         int `json:"next_update"`
	TotalPlayers       int `json:"total_players"`
	RadiantThreshold   int `json:"radiant_threshold"`
	Immortal3Threshold int `json:"immortal_3_threshold"`
	Immortal2Threshold int `json:"immortal_2_threshold"`
	Immortal1Threshold int `json:"immortal_1_threshold"`
	Players            []struct {
		PlayerCardID    string `json:"PlayerCardID"`
		TitleID         string `json:"TitleID"`
		IsBanned        bool   `json:"IsBanned"`
		IsAnonymized    bool   `json:"IsAnonymized"`
		Puuid           string `json:"puuid"`
		GameName        string `json:"gameName"`
		TagLine         string `json:"tagLine"`
		LeaderboardRank int    `json:"leaderboardRank"`
		RankedRating    int    `json:"rankedRating"`
		NumberOfWins    int    `json:"numberOfWins"`
		CompetitiveTier int    `json:"competitiveTier"`
	} `json:"players"`
	Errors []Error `json:"errors"`
}

type GetLifetimeMatchesByNameParams struct {
	Affinity string
	Name     string
	Tag      string
	Mode     string
	Map      string
	Page     string
	Size     string
}

type GetLifetimeMatchesByNameResponse struct {
	Status  int    `json:"status"`
	Name    string `json:"name"`
	Tag     string `json:"tag"`
	Results struct {
		Total    int `json:"total"`
		Returned int `json:"returned"`
		Before   int `json:"before"`
		After    int `json:"after"`
	} `json:"results"`
	Data []struct {
		Meta struct {
			ID  string `json:"id"`
			Map struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"map"`
			Version   string `json:"version"`
			Mode      string `json:"mode"`
			StartedAt string `json:"started_at"`
			Season    struct {
				ID    string `json:"id"`
				Short string `json:"short"`
			} `json:"season"`
			Region  string `json:"region"`
			Cluster string `json:"cluster"`
		} `json:"meta"`
		Stats struct {
			Puuid     string `json:"puuid"`
			Team      string `json:"team"`
			Level     int    `json:"level"`
			Character struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"character"`
			Tier    int `json:"tier"`
			Score   int `json:"score"`
			Kills   int `json:"kills"`
			Deaths  int `json:"deaths"`
			Assists int `json:"assists"`
			Shots   struct {
				Head int `json:"head"`
				Body int `json:"body"`
				Leg  int `json:"leg"`
			} `json:"shots"`
			Damage struct {
				Made     int `json:"made"`
				Received int `json:"received"`
			} `json:"damage"`
		} `json:"stats"`
		Teams struct {
			Red  int `json:"red"`
			Blue int `json:"blue"`
		} `json:"teams"`
	} `json:"data"`
	Errors []Error `json:"errors"`
}

type GetLifetimeMMRHistoryByNameParams struct {
	Affinity string
	Name     string
	Tag      string
	Page     string
	Size     string
}

type GetLifetimeMMRHistoryByNameResponse struct {
	Status  int    `json:"status"`
	Name    string `json:"name"`
	Tag     string `json:"tag"`
	Results struct {
		Total    int `json:"total"`
		Returned int `json:"returned"`
		Before   int `json:"before"`
		After    int `json:"after"`
	} `json:"results"`
	Data []struct {
		MatchID string `json:"match_id"`
		Tier    struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"tier"`
		Map struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"map"`
		Season struct {
			ID    string `json:"id"`
			Short string `json:"short"`
		} `json:"season"`
		RankingInTier int       `json:"ranking_in_tier"`
		LastMmrChange int       `json:"last_mmr_change"`
		Elo           int       `json:"elo"`
		Date          time.Time `json:"date"`
	} `json:"data"`
	Errors []Error `json:"errors"`
}

type GetMatchesByNameV3Params struct {
	Affinity string
	Name     string
	Tag      string
}

type GetMatchesByNameV3Response struct {
	Status int `json:"status"`
	Data   []struct {
		Metadata struct {
			Map              string `json:"map"`
			GameVersion      string `json:"game_version"`
			GameLength       int    `json:"game_length"`
			GameStart        int    `json:"game_start"`
			GameStartPatched string `json:"game_start_patched"`
			RoundsPlayed     int    `json:"rounds_played"`
			Mode             string `json:"mode"`
			ModeID           string `json:"mode_id"`
			Queue            string `json:"queue"`
			SeasonID         string `json:"season_id"`
			Platform         string `json:"platform"`
			Matchid          string `json:"matchid"`
			PremierInfo      struct {
				TournamentID string `json:"tournament_id"`
				MatchupID    string `json:"matchup_id"`
			} `json:"premier_info"`
			Region  string `json:"region"`
			Cluster string `json:"cluster"`
		} `json:"metadata"`
		Players struct {
			AllPlayers []struct {
				Puuid              string `json:"puuid"`
				Name               string `json:"name"`
				Tag                string `json:"tag"`
				Team               string `json:"team"`
				Level              int    `json:"level"`
				Character          string `json:"character"`
				Currenttier        int    `json:"currenttier"`
				CurrenttierPatched string `json:"currenttier_patched"`
				PlayerCard         string `json:"player_card"`
				PlayerTitle        string `json:"player_title"`
				PartyID            string `json:"party_id"`
				SessionPlaytime    struct {
					Minutes      int `json:"minutes"`
					Seconds      int `json:"seconds"`
					Milliseconds int `json:"milliseconds"`
				} `json:"session_playtime"`
				Assets struct {
					Card struct {
						Small string `json:"small"`
						Large string `json:"large"`
						Wide  string `json:"wide"`
					} `json:"card"`
					Agent struct {
						Small    string `json:"small"`
						Full     string `json:"full"`
						Bust     string `json:"bust"`
						Killfeed string `json:"killfeed"`
					} `json:"agent"`
				} `json:"assets"`
				Behaviour struct {
					AfkRounds    int `json:"afk_rounds"`
					FriendlyFire struct {
						Incoming int `json:"incoming"`
						Outgoing int `json:"outgoing"`
					} `json:"friendly_fire"`
					RoundsInSpawn int `json:"rounds_in_spawn"`
				} `json:"behaviour"`
				Platform struct {
					Type string `json:"type"`
					Os   struct {
						Name    string `json:"name"`
						Version string `json:"version"`
					} `json:"os"`
				} `json:"platform"`
				AbilityCasts struct {
					CCast int `json:"c_cast"`
					QCast int `json:"q_cast"`
					ECast int `json:"e_cast"`
					XCast int `json:"x_cast"`
				} `json:"ability_casts"`
				Stats struct {
					Score     int `json:"score"`
					Kills     int `json:"kills"`
					Deaths    int `json:"deaths"`
					Assists   int `json:"assists"`
					Bodyshots int `json:"bodyshots"`
					Headshots int `json:"headshots"`
					Legshots  int `json:"legshots"`
				} `json:"stats"`
				Economy struct {
					Spent struct {
						Overall int `json:"overall"`
						Average int `json:"average"`
					} `json:"spent"`
					LoadoutValue struct {
						Overall int `json:"overall"`
						Average int `json:"average"`
					} `json:"loadout_value"`
				} `json:"economy"`
				DamageMade     int `json:"damage_made"`
				DamageReceived int `json:"damage_received"`
			} `json:"all_players"`
			Red []struct {
				Puuid              string `json:"puuid"`
				Name               string `json:"name"`
				Tag                string `json:"tag"`
				Team               string `json:"team"`
				Level              int    `json:"level"`
				Character          string `json:"character"`
				Currenttier        int    `json:"currenttier"`
				CurrenttierPatched string `json:"currenttier_patched"`
				PlayerCard         string `json:"player_card"`
				PlayerTitle        string `json:"player_title"`
				PartyID            string `json:"party_id"`
				SessionPlaytime    struct {
					Minutes      int `json:"minutes"`
					Seconds      int `json:"seconds"`
					Milliseconds int `json:"milliseconds"`
				} `json:"session_playtime"`
				Assets struct {
					Card struct {
						Small string `json:"small"`
						Large string `json:"large"`
						Wide  string `json:"wide"`
					} `json:"card"`
					Agent struct {
						Small    string `json:"small"`
						Full     string `json:"full"`
						Bust     string `json:"bust"`
						Killfeed string `json:"killfeed"`
					} `json:"agent"`
				} `json:"assets"`
				Behaviour struct {
					AfkRounds    int `json:"afk_rounds"`
					FriendlyFire struct {
						Incoming int `json:"incoming"`
						Outgoing int `json:"outgoing"`
					} `json:"friendly_fire"`
					RoundsInSpawn int `json:"rounds_in_spawn"`
				} `json:"behaviour"`
				Platform struct {
					Type string `json:"type"`
					Os   struct {
						Name    string `json:"name"`
						Version string `json:"version"`
					} `json:"os"`
				} `json:"platform"`
				AbilityCasts struct {
					CCast int `json:"c_cast"`
					QCast int `json:"q_cast"`
					ECast int `json:"e_cast"`
					XCast int `json:"x_cast"`
				} `json:"ability_casts"`
				Stats struct {
					Score     int `json:"score"`
					Kills     int `json:"kills"`
					Deaths    int `json:"deaths"`
					Assists   int `json:"assists"`
					Bodyshots int `json:"bodyshots"`
					Headshots int `json:"headshots"`
					Legshots  int `json:"legshots"`
				} `json:"stats"`
				Economy struct {
					Spent struct {
						Overall int `json:"overall"`
						Average int `json:"average"`
					} `json:"spent"`
					LoadoutValue struct {
						Overall int `json:"overall"`
						Average int `json:"average"`
					} `json:"loadout_value"`
				} `json:"economy"`
				DamageMade     int `json:"damage_made"`
				DamageReceived int `json:"damage_received"`
			} `json:"red"`
			Blue []struct {
				Puuid              string `json:"puuid"`
				Name               string `json:"name"`
				Tag                string `json:"tag"`
				Team               string `json:"team"`
				Level              int    `json:"level"`
				Character          string `json:"character"`
				Currenttier        int    `json:"currenttier"`
				CurrenttierPatched string `json:"currenttier_patched"`
				PlayerCard         string `json:"player_card"`
				PlayerTitle        string `json:"player_title"`
				PartyID            string `json:"party_id"`
				SessionPlaytime    struct {
					Minutes      int `json:"minutes"`
					Seconds      int `json:"seconds"`
					Milliseconds int `json:"milliseconds"`
				} `json:"session_playtime"`
				Assets struct {
					Card struct {
						Small string `json:"small"`
						Large string `json:"large"`
						Wide  string `json:"wide"`
					} `json:"card"`
					Agent struct {
						Small    string `json:"small"`
						Full     string `json:"full"`
						Bust     string `json:"bust"`
						Killfeed string `json:"killfeed"`
					} `json:"agent"`
				} `json:"assets"`
				Behaviour struct {
					AfkRounds    int `json:"afk_rounds"`
					FriendlyFire struct {
						Incoming int `json:"incoming"`
						Outgoing int `json:"outgoing"`
					} `json:"friendly_fire"`
					RoundsInSpawn int `json:"rounds_in_spawn"`
				} `json:"behaviour"`
				Platform struct {
					Type string `json:"type"`
					Os   struct {
						Name    string `json:"name"`
						Version string `json:"version"`
					} `json:"os"`
				} `json:"platform"`
				AbilityCasts struct {
					CCast int `json:"c_cast"`
					QCast int `json:"q_cast"`
					ECast int `json:"e_cast"`
					XCast int `json:"x_cast"`
				} `json:"ability_casts"`
				Stats struct {
					Score     int `json:"score"`
					Kills     int `json:"kills"`
					Deaths    int `json:"deaths"`
					Assists   int `json:"assists"`
					Bodyshots int `json:"bodyshots"`
					Headshots int `json:"headshots"`
					Legshots  int `json:"legshots"`
				} `json:"stats"`
				Economy struct {
					Spent struct {
						Overall int `json:"overall"`
						Average int `json:"average"`
					} `json:"spent"`
					LoadoutValue struct {
						Overall int `json:"overall"`
						Average int `json:"average"`
					} `json:"loadout_value"`
				} `json:"economy"`
				DamageMade     int `json:"damage_made"`
				DamageReceived int `json:"damage_received"`
			} `json:"blue"`
		} `json:"players"`
		Observers []struct {
			Puuid    string `json:"puuid"`
			Name     string `json:"name"`
			Tag      string `json:"tag"`
			Platform struct {
				Type string `json:"type"`
				Os   struct {
					Name    string `json:"name"`
					Version string `json:"version"`
				} `json:"os"`
			} `json:"platform"`
			SessionPlaytime struct {
				Minutes      int `json:"minutes"`
				Seconds      int `json:"seconds"`
				Milliseconds int `json:"milliseconds"`
			} `json:"session_playtime"`
			Team        string `json:"team"`
			Level       int    `json:"level"`
			PlayerCard  string `json:"player_card"`
			PlayerTitle string `json:"player_title"`
			PartyID     string `json:"party_id"`
		} `json:"observers"`
		Coaches []struct {
			Puuid string `json:"puuid"`
			Team  string `json:"team"`
		} `json:"coaches"`
		Teams struct {
			Red struct {
				HasWon     bool `json:"has_won"`
				RoundsWon  int  `json:"rounds_won"`
				RoundsLost int  `json:"rounds_lost"`
				Roaster    struct {
					Members       []string `json:"members"`
					Name          string   `json:"name"`
					Tag           string   `json:"tag"`
					Customization struct {
						Icon      string `json:"icon"`
						Image     string `json:"image"`
						Primary   string `json:"primary"`
						Secondary string `json:"secondary"`
						Tertiary  string `json:"tertiary"`
					} `json:"customization"`
				} `json:"roaster"`
			} `json:"red"`
			Blue struct {
				HasWon     bool `json:"has_won"`
				RoundsWon  int  `json:"rounds_won"`
				RoundsLost int  `json:"rounds_lost"`
				Roaster    struct {
					Members       []string `json:"members"`
					Name          string   `json:"name"`
					Tag           string   `json:"tag"`
					Customization struct {
						Icon      string `json:"icon"`
						Image     string `json:"image"`
						Primary   string `json:"primary"`
						Secondary string `json:"secondary"`
						Tertiary  string `json:"tertiary"`
					} `json:"customization"`
				} `json:"roaster"`
			} `json:"blue"`
		} `json:"teams"`
		Rounds []struct {
			WinningTeam string `json:"winning_team"`
			EndType     string `json:"end_type"`
			BombPlanted bool   `json:"bomb_planted"`
			BombDefused bool   `json:"bomb_defused"`
			PlantEvents struct {
				PlantLocation struct {
					X int `json:"x"`
					Y int `json:"y"`
				} `json:"plant_location"`
				PlantedBy struct {
					Puuid       string `json:"puuid"`
					DisplayName string `json:"display_name"`
					Team        string `json:"team"`
				} `json:"planted_by"`
				PlantSite              string `json:"plant_site"`
				PlantTimeInRound       int    `json:"plant_time_in_round"`
				PlayerLocationsOnPlant []struct {
					PlayerPuuid       string `json:"player_puuid"`
					PlayerDisplayName string `json:"player_display_name"`
					PlayerTeam        string `json:"player_team"`
					Location          struct {
						X int `json:"x"`
						Y int `json:"y"`
					} `json:"location"`
					ViewRadians float64 `json:"view_radians"`
				} `json:"player_locations_on_plant"`
			} `json:"plant_events"`
			DefuseEvents struct {
				DefuseLocation struct {
					X int `json:"x"`
					Y int `json:"y"`
				} `json:"defuse_location"`
				DefusedBy struct {
					Puuid       string `json:"puuid"`
					DisplayName string `json:"display_name"`
					Team        string `json:"team"`
				} `json:"defused_by"`
				DefuseTimeInRound       int `json:"defuse_time_in_round"`
				PlayerLocationsOnDefuse []struct {
					PlayerPuuid       string `json:"player_puuid"`
					PlayerDisplayName string `json:"player_display_name"`
					PlayerTeam        string `json:"player_team"`
					Location          struct {
						X int `json:"x"`
						Y int `json:"y"`
					} `json:"location"`
					ViewRadians float64 `json:"view_radians"`
				} `json:"player_locations_on_defuse"`
			} `json:"defuse_events"`
			PlayerStats []struct {
				AbilityCasts struct {
					CCasts int `json:"c_casts"`
					QCasts int `json:"q_casts"`
					ECasts int `json:"e_casts"`
					XCasts int `json:"x_casts"`
				} `json:"ability_casts"`
				PlayerPuuid       string        `json:"player_puuid"`
				PlayerDisplayName string        `json:"player_display_name"`
				PlayerTeam        string        `json:"player_team"`
				DamageEvents      []interface{} `json:"damage_events"`
				Damage            int           `json:"damage"`
				Bodyshots         int           `json:"bodyshots"`
				Headshots         int           `json:"headshots"`
				Legshots          int           `json:"legshots"`
				KillEvents        []interface{} `json:"kill_events"`
				Kills             int           `json:"kills"`
				Score             int           `json:"score"`
				Economy           struct {
					LoadoutValue int `json:"loadout_value"`
					Weapon       struct {
						ID     string `json:"id"`
						Name   string `json:"name"`
						Assets struct {
							DisplayIcon  string `json:"display_icon"`
							KillfeedIcon string `json:"killfeed_icon"`
						} `json:"assets"`
					} `json:"weapon"`
					Armor struct {
						ID     string `json:"id"`
						Name   string `json:"name"`
						Assets struct {
							DisplayIcon string `json:"display_icon"`
						} `json:"assets"`
					} `json:"armor"`
					Remaining int `json:"remaining"`
					Spent     int `json:"spent"`
				} `json:"economy"`
				WasAfk        bool `json:"was_afk"`
				WasPenalized  bool `json:"was_penalized"`
				StayedInSpawn bool `json:"stayed_in_spawn"`
			} `json:"player_stats"`
		} `json:"rounds"`
	} `json:"data"`
	Errors []Error `json:"errors"`
}

type GetMatchParams struct {
	MatchId string
}

type GetMatchResponse struct {
	Status int `json:"status"`
	Data   struct {
		Metadata struct {
			Map              string `json:"map"`
			GameVersion      string `json:"game_version"`
			GameLength       int    `json:"game_length"`
			GameStart        int    `json:"game_start"`
			GameStartPatched string `json:"game_start_patched"`
			RoundsPlayed     int    `json:"rounds_played"`
			Mode             string `json:"mode"`
			ModeID           string `json:"mode_id"`
			Queue            string `json:"queue"`
			SeasonID         string `json:"season_id"`
			Platform         string `json:"platform"`
			Matchid          string `json:"matchid"`
			PremierInfo      struct {
				TournamentID string `json:"tournament_id"`
				MatchupID    string `json:"matchup_id"`
			} `json:"premier_info"`
			Region  string `json:"region"`
			Cluster string `json:"cluster"`
		} `json:"metadata"`
		Players struct {
			AllPlayers []struct {
				Puuid              string `json:"puuid"`
				Name               string `json:"name"`
				Tag                string `json:"tag"`
				Team               string `json:"team"`
				Level              int    `json:"level"`
				Character          string `json:"character"`
				Currenttier        int    `json:"currenttier"`
				CurrenttierPatched string `json:"currenttier_patched"`
				PlayerCard         string `json:"player_card"`
				PlayerTitle        string `json:"player_title"`
				PartyID            string `json:"party_id"`
				SessionPlaytime    struct {
					Minutes      int `json:"minutes"`
					Seconds      int `json:"seconds"`
					Milliseconds int `json:"milliseconds"`
				} `json:"session_playtime"`
				Assets struct {
					Card struct {
						Small string `json:"small"`
						Large string `json:"large"`
						Wide  string `json:"wide"`
					} `json:"card"`
					Agent struct {
						Small    string `json:"small"`
						Full     string `json:"full"`
						Bust     string `json:"bust"`
						Killfeed string `json:"killfeed"`
					} `json:"agent"`
				} `json:"assets"`
				Behaviour struct {
					AfkRounds    int `json:"afk_rounds"`
					FriendlyFire struct {
						Incoming int `json:"incoming"`
						Outgoing int `json:"outgoing"`
					} `json:"friendly_fire"`
					RoundsInSpawn int `json:"rounds_in_spawn"`
				} `json:"behaviour"`
				Platform struct {
					Type string `json:"type"`
					Os   struct {
						Name    string `json:"name"`
						Version string `json:"version"`
					} `json:"os"`
				} `json:"platform"`
				AbilityCasts struct {
					CCast int `json:"c_cast"`
					QCast int `json:"q_cast"`
					ECast int `json:"e_cast"`
					XCast int `json:"x_cast"`
				} `json:"ability_casts"`
				Stats struct {
					Score     int `json:"score"`
					Kills     int `json:"kills"`
					Deaths    int `json:"deaths"`
					Assists   int `json:"assists"`
					Bodyshots int `json:"bodyshots"`
					Headshots int `json:"headshots"`
					Legshots  int `json:"legshots"`
				} `json:"stats"`
				Economy struct {
					Spent struct {
						Overall int `json:"overall"`
						Average int `json:"average"`
					} `json:"spent"`
					LoadoutValue struct {
						Overall int `json:"overall"`
						Average int `json:"average"`
					} `json:"loadout_value"`
				} `json:"economy"`
				DamageMade     int `json:"damage_made"`
				DamageReceived int `json:"damage_received"`
			} `json:"all_players"`
			Red []struct {
				Puuid              string `json:"puuid"`
				Name               string `json:"name"`
				Tag                string `json:"tag"`
				Team               string `json:"team"`
				Level              int    `json:"level"`
				Character          string `json:"character"`
				Currenttier        int    `json:"currenttier"`
				CurrenttierPatched string `json:"currenttier_patched"`
				PlayerCard         string `json:"player_card"`
				PlayerTitle        string `json:"player_title"`
				PartyID            string `json:"party_id"`
				SessionPlaytime    struct {
					Minutes      int `json:"minutes"`
					Seconds      int `json:"seconds"`
					Milliseconds int `json:"milliseconds"`
				} `json:"session_playtime"`
				Assets struct {
					Card struct {
						Small string `json:"small"`
						Large string `json:"large"`
						Wide  string `json:"wide"`
					} `json:"card"`
					Agent struct {
						Small    string `json:"small"`
						Full     string `json:"full"`
						Bust     string `json:"bust"`
						Killfeed string `json:"killfeed"`
					} `json:"agent"`
				} `json:"assets"`
				Behaviour struct {
					AfkRounds    int `json:"afk_rounds"`
					FriendlyFire struct {
						Incoming int `json:"incoming"`
						Outgoing int `json:"outgoing"`
					} `json:"friendly_fire"`
					RoundsInSpawn int `json:"rounds_in_spawn"`
				} `json:"behaviour"`
				Platform struct {
					Type string `json:"type"`
					Os   struct {
						Name    string `json:"name"`
						Version string `json:"version"`
					} `json:"os"`
				} `json:"platform"`
				AbilityCasts struct {
					CCast int `json:"c_cast"`
					QCast int `json:"q_cast"`
					ECast int `json:"e_cast"`
					XCast int `json:"x_cast"`
				} `json:"ability_casts"`
				Stats struct {
					Score     int `json:"score"`
					Kills     int `json:"kills"`
					Deaths    int `json:"deaths"`
					Assists   int `json:"assists"`
					Bodyshots int `json:"bodyshots"`
					Headshots int `json:"headshots"`
					Legshots  int `json:"legshots"`
				} `json:"stats"`
				Economy struct {
					Spent struct {
						Overall int `json:"overall"`
						Average int `json:"average"`
					} `json:"spent"`
					LoadoutValue struct {
						Overall int `json:"overall"`
						Average int `json:"average"`
					} `json:"loadout_value"`
				} `json:"economy"`
				DamageMade     int `json:"damage_made"`
				DamageReceived int `json:"damage_received"`
			} `json:"red"`
			Blue []struct {
				Puuid              string `json:"puuid"`
				Name               string `json:"name"`
				Tag                string `json:"tag"`
				Team               string `json:"team"`
				Level              int    `json:"level"`
				Character          string `json:"character"`
				Currenttier        int    `json:"currenttier"`
				CurrenttierPatched string `json:"currenttier_patched"`
				PlayerCard         string `json:"player_card"`
				PlayerTitle        string `json:"player_title"`
				PartyID            string `json:"party_id"`
				SessionPlaytime    struct {
					Minutes      int `json:"minutes"`
					Seconds      int `json:"seconds"`
					Milliseconds int `json:"milliseconds"`
				} `json:"session_playtime"`
				Assets struct {
					Card struct {
						Small string `json:"small"`
						Large string `json:"large"`
						Wide  string `json:"wide"`
					} `json:"card"`
					Agent struct {
						Small    string `json:"small"`
						Full     string `json:"full"`
						Bust     string `json:"bust"`
						Killfeed string `json:"killfeed"`
					} `json:"agent"`
				} `json:"assets"`
				Behaviour struct {
					AfkRounds    int `json:"afk_rounds"`
					FriendlyFire struct {
						Incoming int `json:"incoming"`
						Outgoing int `json:"outgoing"`
					} `json:"friendly_fire"`
					RoundsInSpawn int `json:"rounds_in_spawn"`
				} `json:"behaviour"`
				Platform struct {
					Type string `json:"type"`
					Os   struct {
						Name    string `json:"name"`
						Version string `json:"version"`
					} `json:"os"`
				} `json:"platform"`
				AbilityCasts struct {
					CCast int `json:"c_cast"`
					QCast int `json:"q_cast"`
					ECast int `json:"e_cast"`
					XCast int `json:"x_cast"`
				} `json:"ability_casts"`
				Stats struct {
					Score     int `json:"score"`
					Kills     int `json:"kills"`
					Deaths    int `json:"deaths"`
					Assists   int `json:"assists"`
					Bodyshots int `json:"bodyshots"`
					Headshots int `json:"headshots"`
					Legshots  int `json:"legshots"`
				} `json:"stats"`
				Economy struct {
					Spent struct {
						Overall int `json:"overall"`
						Average int `json:"average"`
					} `json:"spent"`
					LoadoutValue struct {
						Overall int `json:"overall"`
						Average int `json:"average"`
					} `json:"loadout_value"`
				} `json:"economy"`
				DamageMade     int `json:"damage_made"`
				DamageReceived int `json:"damage_received"`
			} `json:"blue"`
		} `json:"players"`
		Observers []struct {
			Puuid    string `json:"puuid"`
			Name     string `json:"name"`
			Tag      string `json:"tag"`
			Platform struct {
				Type string `json:"type"`
				Os   struct {
					Name    string `json:"name"`
					Version string `json:"version"`
				} `json:"os"`
			} `json:"platform"`
			SessionPlaytime struct {
				Minutes      int `json:"minutes"`
				Seconds      int `json:"seconds"`
				Milliseconds int `json:"milliseconds"`
			} `json:"session_playtime"`
			Team        string `json:"team"`
			Level       int    `json:"level"`
			PlayerCard  string `json:"player_card"`
			PlayerTitle string `json:"player_title"`
			PartyID     string `json:"party_id"`
		} `json:"observers"`
		Coaches []struct {
			Puuid string `json:"puuid"`
			Team  string `json:"team"`
		} `json:"coaches"`
		Teams struct {
			Red struct {
				HasWon     bool `json:"has_won"`
				RoundsWon  int  `json:"rounds_won"`
				RoundsLost int  `json:"rounds_lost"`
				Roaster    struct {
					Members       []string `json:"members"`
					Name          string   `json:"name"`
					Tag           string   `json:"tag"`
					Customization struct {
						Icon      string `json:"icon"`
						Image     string `json:"image"`
						Primary   string `json:"primary"`
						Secondary string `json:"secondary"`
						Tertiary  string `json:"tertiary"`
					} `json:"customization"`
				} `json:"roaster"`
			} `json:"red"`
			Blue struct {
				HasWon     bool `json:"has_won"`
				RoundsWon  int  `json:"rounds_won"`
				RoundsLost int  `json:"rounds_lost"`
				Roaster    struct {
					Members       []string `json:"members"`
					Name          string   `json:"name"`
					Tag           string   `json:"tag"`
					Customization struct {
						Icon      string `json:"icon"`
						Image     string `json:"image"`
						Primary   string `json:"primary"`
						Secondary string `json:"secondary"`
						Tertiary  string `json:"tertiary"`
					} `json:"customization"`
				} `json:"roaster"`
			} `json:"blue"`
		} `json:"teams"`
		Rounds []struct {
			WinningTeam string `json:"winning_team"`
			EndType     string `json:"end_type"`
			BombPlanted bool   `json:"bomb_planted"`
			BombDefused bool   `json:"bomb_defused"`
			PlantEvents struct {
				PlantLocation struct {
					X int `json:"x"`
					Y int `json:"y"`
				} `json:"plant_location"`
				PlantedBy struct {
					Puuid       string `json:"puuid"`
					DisplayName string `json:"display_name"`
					Team        string `json:"team"`
				} `json:"planted_by"`
				PlantSite              string `json:"plant_site"`
				PlantTimeInRound       int    `json:"plant_time_in_round"`
				PlayerLocationsOnPlant []struct {
					PlayerPuuid       string `json:"player_puuid"`
					PlayerDisplayName string `json:"player_display_name"`
					PlayerTeam        string `json:"player_team"`
					Location          struct {
						X int `json:"x"`
						Y int `json:"y"`
					} `json:"location"`
					ViewRadians float64 `json:"view_radians"`
				} `json:"player_locations_on_plant"`
			} `json:"plant_events"`
			DefuseEvents struct {
				DefuseLocation struct {
					X int `json:"x"`
					Y int `json:"y"`
				} `json:"defuse_location"`
				DefusedBy struct {
					Puuid       string `json:"puuid"`
					DisplayName string `json:"display_name"`
					Team        string `json:"team"`
				} `json:"defused_by"`
				DefuseTimeInRound       int `json:"defuse_time_in_round"`
				PlayerLocationsOnDefuse []struct {
					PlayerPuuid       string `json:"player_puuid"`
					PlayerDisplayName string `json:"player_display_name"`
					PlayerTeam        string `json:"player_team"`
					Location          struct {
						X int `json:"x"`
						Y int `json:"y"`
					} `json:"location"`
					ViewRadians float64 `json:"view_radians"`
				} `json:"player_locations_on_defuse"`
			} `json:"defuse_events"`
			PlayerStats []struct {
				AbilityCasts struct {
					CCasts int `json:"c_casts"`
					QCasts int `json:"q_casts"`
					ECasts int `json:"e_casts"`
					XCasts int `json:"x_casts"`
				} `json:"ability_casts"`
				PlayerPuuid       string        `json:"player_puuid"`
				PlayerDisplayName string        `json:"player_display_name"`
				PlayerTeam        string        `json:"player_team"`
				DamageEvents      []interface{} `json:"damage_events"`
				Damage            int           `json:"damage"`
				Bodyshots         int           `json:"bodyshots"`
				Headshots         int           `json:"headshots"`
				Legshots          int           `json:"legshots"`
				KillEvents        []interface{} `json:"kill_events"`
				Kills             int           `json:"kills"`
				Score             int           `json:"score"`
				Economy           struct {
					LoadoutValue int `json:"loadout_value"`
					Weapon       struct {
						ID     string `json:"id"`
						Name   string `json:"name"`
						Assets struct {
							DisplayIcon  string `json:"display_icon"`
							KillfeedIcon string `json:"killfeed_icon"`
						} `json:"assets"`
					} `json:"weapon"`
					Armor struct {
						ID     string `json:"id"`
						Name   string `json:"name"`
						Assets struct {
							DisplayIcon string `json:"display_icon"`
						} `json:"assets"`
					} `json:"armor"`
					Remaining int `json:"remaining"`
					Spent     int `json:"spent"`
				} `json:"economy"`
				WasAfk        bool `json:"was_afk"`
				WasPenalized  bool `json:"was_penalized"`
				StayedInSpawn bool `json:"stayed_in_spawn"`
			} `json:"player_stats"`
		} `json:"rounds"`
	} `json:"data"`
	Errors []Error `json:"errors"`
}

type GetMMRHistoryByNameParams struct {
	Affinity string
	Name     string
	Tag      string
}

type GetMMRHistoryByNameResponse struct {
	Status int    `json:"status"`
	Name   string `json:"name"`
	Tag    string `json:"tag"`
	Data   []struct {
		Currenttier        int    `json:"currenttier"`
		CurrenttierPatched string `json:"currenttier_patched"`
		Images             struct {
			Small        string `json:"small"`
			Large        string `json:"large"`
			TriangleDown string `json:"triangle_down"`
			TriangleUp   string `json:"triangle_up"`
		} `json:"images"`
		MatchID string `json:"match_id"`
		Map     struct {
			Name string `json:"name"`
			ID   string `json:"id"`
		} `json:"map"`
		SeasonID            string `json:"season_id"`
		RankingInTier       int    `json:"ranking_in_tier"`
		MmrChangeToLastGame int    `json:"mmr_change_to_last_game"`
		Elo                 int    `json:"elo"`
		Date                string `json:"date"`
		DateRaw             int    `json:"date_raw"`
	} `json:"data"`
	Errors []Error `json:"errors"`
}

type GetMMRByNameV2Params struct {
	Affinity string
	Name     string
	Tag      string
	Season   string
}

type GetMMRByNameV2Response struct {
	Status int `json:"status"`
	Data   struct {
		Name        string `json:"name"`
		Tag         string `json:"tag"`
		CurrentData struct {
			Currenttier        int    `json:"currenttier"`
			CurrenttierPatched string `json:"currenttier_patched"`
			Images             struct {
				Small        string `json:"small"`
				Large        string `json:"large"`
				TriangleDown string `json:"triangle_down"`
				TriangleUp   string `json:"triangle_up"`
			} `json:"images"`
			RankingInTier       int  `json:"ranking_in_tier"`
			MmrChangeToLastGame int  `json:"mmr_change_to_last_game"`
			Elo                 int  `json:"elo"`
			Old                 bool `json:"old"`
		} `json:"current_data"`
		HighestRank struct {
			Old         bool   `json:"old"`
			Tier        int    `json:"tier"`
			PatchedTier string `json:"patched_tier"`
			Season      string `json:"season"`
		} `json:"highest_rank"`
		BySeason struct {
			E6A3 struct {
				Error            string `json:"error"`
				Wins             int    `json:"wins"`
				NumberOfGames    int    `json:"number_of_games"`
				FinalRank        int    `json:"final_rank"`
				FinalRankPatched string `json:"final_rank_patched"`
				ActRankWins      []struct {
					PatchedTier string `json:"patched_tier"`
					Tier        int    `json:"tier"`
				} `json:"act_rank_wins"`
				Old bool `json:"old"`
			} `json:"e6a3"`
			E6A2 struct {
				Error            string `json:"error"`
				Wins             int    `json:"wins"`
				NumberOfGames    int    `json:"number_of_games"`
				FinalRank        int    `json:"final_rank"`
				FinalRankPatched string `json:"final_rank_patched"`
				ActRankWins      []struct {
					PatchedTier string `json:"patched_tier"`
					Tier        int    `json:"tier"`
				} `json:"act_rank_wins"`
				Old bool `json:"old"`
			} `json:"e6a2"`
			E6A1 struct {
				Error            string `json:"error"`
				Wins             int    `json:"wins"`
				NumberOfGames    int    `json:"number_of_games"`
				FinalRank        int    `json:"final_rank"`
				FinalRankPatched string `json:"final_rank_patched"`
				ActRankWins      []struct {
					PatchedTier string `json:"patched_tier"`
					Tier        int    `json:"tier"`
				} `json:"act_rank_wins"`
				Old bool `json:"old"`
			} `json:"e6a1"`
			E5A3 struct {
				Error            string `json:"error"`
				Wins             int    `json:"wins"`
				NumberOfGames    int    `json:"number_of_games"`
				FinalRank        int    `json:"final_rank"`
				FinalRankPatched string `json:"final_rank_patched"`
				ActRankWins      []struct {
					PatchedTier string `json:"patched_tier"`
					Tier        int    `json:"tier"`
				} `json:"act_rank_wins"`
				Old bool `json:"old"`
			} `json:"e5a3"`
			E5A2 struct {
				Error            string `json:"error"`
				Wins             int    `json:"wins"`
				NumberOfGames    int    `json:"number_of_games"`
				FinalRank        int    `json:"final_rank"`
				FinalRankPatched string `json:"final_rank_patched"`
				ActRankWins      []struct {
					PatchedTier string `json:"patched_tier"`
					Tier        int    `json:"tier"`
				} `json:"act_rank_wins"`
				Old bool `json:"old"`
			} `json:"e5a2"`
			E5A1 struct {
				Error            string `json:"error"`
				Wins             int    `json:"wins"`
				NumberOfGames    int    `json:"number_of_games"`
				FinalRank        int    `json:"final_rank"`
				FinalRankPatched string `json:"final_rank_patched"`
				ActRankWins      []struct {
					PatchedTier string `json:"patched_tier"`
					Tier        int    `json:"tier"`
				} `json:"act_rank_wins"`
				Old bool `json:"old"`
			} `json:"e5a1"`
			E4A3 struct {
				Error            string `json:"error"`
				Wins             int    `json:"wins"`
				NumberOfGames    int    `json:"number_of_games"`
				FinalRank        int    `json:"final_rank"`
				FinalRankPatched string `json:"final_rank_patched"`
				ActRankWins      []struct {
					PatchedTier string `json:"patched_tier"`
					Tier        int    `json:"tier"`
				} `json:"act_rank_wins"`
				Old bool `json:"old"`
			} `json:"e4a3"`
			E4A2 struct {
				Error            string `json:"error"`
				Wins             int    `json:"wins"`
				NumberOfGames    int    `json:"number_of_games"`
				FinalRank        int    `json:"final_rank"`
				FinalRankPatched string `json:"final_rank_patched"`
				ActRankWins      []struct {
					PatchedTier string `json:"patched_tier"`
					Tier        int    `json:"tier"`
				} `json:"act_rank_wins"`
				Old bool `json:"old"`
			} `json:"e4a2"`
			E4A1 struct {
				Error            string `json:"error"`
				Wins             int    `json:"wins"`
				NumberOfGames    int    `json:"number_of_games"`
				FinalRank        int    `json:"final_rank"`
				FinalRankPatched string `json:"final_rank_patched"`
				ActRankWins      []struct {
					PatchedTier string `json:"patched_tier"`
					Tier        int    `json:"tier"`
				} `json:"act_rank_wins"`
				Old bool `json:"old"`
			} `json:"e4a1"`
			E3A3 struct {
				Error            string `json:"error"`
				Wins             int    `json:"wins"`
				NumberOfGames    int    `json:"number_of_games"`
				FinalRank        int    `json:"final_rank"`
				FinalRankPatched string `json:"final_rank_patched"`
				ActRankWins      []struct {
					PatchedTier string `json:"patched_tier"`
					Tier        int    `json:"tier"`
				} `json:"act_rank_wins"`
				Old bool `json:"old"`
			} `json:"e3a3"`
			E3A2 struct {
				Error            string `json:"error"`
				Wins             int    `json:"wins"`
				NumberOfGames    int    `json:"number_of_games"`
				FinalRank        int    `json:"final_rank"`
				FinalRankPatched string `json:"final_rank_patched"`
				ActRankWins      []struct {
					PatchedTier string `json:"patched_tier"`
					Tier        int    `json:"tier"`
				} `json:"act_rank_wins"`
				Old bool `json:"old"`
			} `json:"e3a2"`
			E3A1 struct {
				Error            string `json:"error"`
				Wins             int    `json:"wins"`
				NumberOfGames    int    `json:"number_of_games"`
				FinalRank        int    `json:"final_rank"`
				FinalRankPatched string `json:"final_rank_patched"`
				ActRankWins      []struct {
					PatchedTier string `json:"patched_tier"`
					Tier        int    `json:"tier"`
				} `json:"act_rank_wins"`
				Old bool `json:"old"`
			} `json:"e3a1"`
			E2A3 struct {
				Error            string `json:"error"`
				Wins             int    `json:"wins"`
				NumberOfGames    int    `json:"number_of_games"`
				FinalRank        int    `json:"final_rank"`
				FinalRankPatched string `json:"final_rank_patched"`
				ActRankWins      []struct {
					PatchedTier string `json:"patched_tier"`
					Tier        int    `json:"tier"`
				} `json:"act_rank_wins"`
				Old bool `json:"old"`
			} `json:"e2a3"`
			E2A2 struct {
				Error            string `json:"error"`
				Wins             int    `json:"wins"`
				NumberOfGames    int    `json:"number_of_games"`
				FinalRank        int    `json:"final_rank"`
				FinalRankPatched string `json:"final_rank_patched"`
				ActRankWins      []struct {
					PatchedTier string `json:"patched_tier"`
					Tier        int    `json:"tier"`
				} `json:"act_rank_wins"`
				Old bool `json:"old"`
			} `json:"e2a2"`
			E2A1 struct {
				Error            string `json:"error"`
				Wins             int    `json:"wins"`
				NumberOfGames    int    `json:"number_of_games"`
				FinalRank        int    `json:"final_rank"`
				FinalRankPatched string `json:"final_rank_patched"`
				ActRankWins      []struct {
					PatchedTier string `json:"patched_tier"`
					Tier        int    `json:"tier"`
				} `json:"act_rank_wins"`
				Old bool `json:"old"`
			} `json:"e2a1"`
			E1A3 struct {
				Error            string `json:"error"`
				Wins             int    `json:"wins"`
				NumberOfGames    int    `json:"number_of_games"`
				FinalRank        int    `json:"final_rank"`
				FinalRankPatched string `json:"final_rank_patched"`
				ActRankWins      []struct {
					PatchedTier string `json:"patched_tier"`
					Tier        int    `json:"tier"`
				} `json:"act_rank_wins"`
				Old bool `json:"old"`
			} `json:"e1a3"`
			E1A2 struct {
				Error            string `json:"error"`
				Wins             int    `json:"wins"`
				NumberOfGames    int    `json:"number_of_games"`
				FinalRank        int    `json:"final_rank"`
				FinalRankPatched string `json:"final_rank_patched"`
				ActRankWins      []struct {
					PatchedTier string `json:"patched_tier"`
					Tier        int    `json:"tier"`
				} `json:"act_rank_wins"`
				Old bool `json:"old"`
			} `json:"e1a2"`
			E1A1 struct {
				Error            string `json:"error"`
				Wins             int    `json:"wins"`
				NumberOfGames    int    `json:"number_of_games"`
				FinalRank        int    `json:"final_rank"`
				FinalRankPatched string `json:"final_rank_patched"`
				ActRankWins      []struct {
					PatchedTier string `json:"patched_tier"`
					Tier        int    `json:"tier"`
				} `json:"act_rank_wins"`
				Old bool `json:"old"`
			} `json:"e1a1"`
		} `json:"by_season"`
		Wins             int    `json:"wins"`
		NumberOfGames    int    `json:"number_of_games"`
		FinalRank        int    `json:"final_rank"`
		FinalRankPatched string `json:"final_rank_patched"`
		ActRankWins      []struct {
			PatchedTier string `json:"patched_tier"`
			Tier        int    `json:"tier"`
		} `json:"act_rank_wins"`
		Old bool `json:"old"`
	} `json:"data"`
	Errors []Error `json:"errors"`
}

type GetPremierTeamParams struct {
	TeamName string
	TeamTag  string
}

type GetPremierTeamResponse struct {
	Status int `json:"status"`
	Data   struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		Tag      string `json:"tag"`
		Enrolled bool   `json:"enrolled"`
		Stats    struct {
			Wins    int `json:"wins"`
			Matches int `json:"matches"`
			Losses  int `json:"losses"`
		} `json:"stats"`
		Placement struct {
			Points     int    `json:"points"`
			Conference string `json:"conference"`
			Division   int    `json:"division"`
			Place      int    `json:"place"`
		} `json:"placement"`
		Customization struct {
			Icon      string `json:"icon"`
			Image     string `json:"image"`
			Primary   string `json:"primary"`
			Secondary string `json:"secondary"`
			Tertiary  string `json:"tertiary"`
		} `json:"customization"`
		Member []struct {
			Puuid string `json:"puuid"`
			Name  string `json:"name"`
			Tag   string `json:"tag"`
		} `json:"member"`
	} `json:"data"`
	Errors []Error `json:"errors"`
}

type GetPremierTeamHistoryResponse struct {
	Status int `json:"status"`
	Data   struct {
		LeagueMatches []struct {
			ID           string    `json:"id"`
			PointsBefore int       `json:"points_before"`
			PointsAfter  int       `json:"points_after"`
			StartedAt    time.Time `json:"started_at"`
		} `json:"league_matches"`
	} `json:"data"`
	Errors []Error `json:"errors"`
}

type GetPremierTeamsParams struct {
	TeamName   string
	TeamTag    string
	Division   string
	Conference string
}

type GetPremierTeamsResponse struct {
	Status int `json:"status"`
	Data   []struct {
		ID            string `json:"id"`
		Name          string `json:"name"`
		Tag           string `json:"tag"`
		Conference    string `json:"conference"`
		Division      int    `json:"division"`
		Affinity      string `json:"affinity"`
		Region        string `json:"region"`
		Losses        int    `json:"losses"`
		Wins          int    `json:"wins"`
		Score         int    `json:"score"`
		Ranking       int    `json:"ranking"`
		Customization struct {
			Icon      string `json:"icon"`
			Image     string `json:"image"`
			Primary   string `json:"primary"`
			Secondary string `json:"secondary"`
			Tertiary  string `json:"tertiary"`
		} `json:"customization"`
	} `json:"data"`
	Errors []Error `json:"errors"`
}

type GetPremierConferencesResponse struct {
	Status int `json:"status"`
	Data   []struct {
		ID       string `json:"id"`
		Affinity string `json:"affinity"`
		Pods     []struct {
			Pod  string `json:"pod"`
			Name string `json:"name"`
		} `json:"pods"`
		Region   string `json:"region"`
		Timezone string `json:"timezone"`
		Name     string `json:"name"`
		Icon     string `json:"icon"`
	} `json:"data"`
	Erros []Error `json:"errors"`
}

type GetPremierSeasonsParams struct {
	Affinity string
}

type GetPremierSeasonsResponse struct {
	Status int `json:"status"`
	Data   []struct {
		ID                         string    `json:"id"`
		ChampionshipEventID        string    `json:"championship_event_id"`
		ChampionshipPointsRequired int       `json:"championship_points_required"`
		StartsAt                   time.Time `json:"starts_at"`
		EndsAt                     time.Time `json:"ends_at"`
		EnrollmentStartsAt         time.Time `json:"enrollment_starts_at"`
		EnrollmentEndsAt           time.Time `json:"enrollment_ends_at"`
		Events                     []struct {
			ID                  string    `json:"id"`
			Type                string    `json:"type"`
			StartsAt            time.Time `json:"starts_at"`
			EndsAt              time.Time `json:"ends_at"`
			ConferenceSchedules []struct {
				Conference string    `json:"conference"`
				StartsAt   time.Time `json:"starts_at"`
				EndsAt     time.Time `json:"ends_at"`
			} `json:"conference_schedules"`
			MapSelection struct {
				Type string `json:"type"`
				Maps []struct {
					Name string `json:"name"`
					ID   string `json:"id"`
				} `json:"maps"`
			} `json:"map_selection"`
			PointsRequiredToParticipate int `json:"points_required_to_participate"`
		} `json:"events"`
	} `json:"data"`
	Errors []Error `json:"errors"`
}

type GetPremierLeaderboardParams struct {
	Affinity string
}

type GetPremierLeaderboardResponse struct {
	Status int `json:"status"`
	Data   []struct {
		ID            string `json:"id"`
		Name          string `json:"name"`
		Tag           string `json:"tag"`
		Conference    string `json:"conference"`
		Division      int    `json:"division"`
		Affinity      string `json:"affinity"`
		Region        string `json:"region"`
		Losses        int    `json:"losses"`
		Wins          int    `json:"wins"`
		Score         int    `json:"score"`
		Ranking       int    `json:"ranking"`
		Customization struct {
			Icon      string `json:"icon"`
			Image     string `json:"image"`
			Primary   string `json:"primary"`
			Secondary string `json:"secondary"`
			Tertiary  string `json:"tertiary"`
		} `json:"customization"`
	} `json:"data"`
	Errors []Error `json:"errors"`
}

type GetPremierConfLeaderboardParams struct {
	Affinity   string
	Conference string
}

type GetPremierConfDivLeaderboardParams struct {
	Affinity   string
	Conference string
	Division   string
}

type GetStatusParams struct {
	Affinity string
}

type GetQueueStatusResponse struct {
	Status int `json:"status"`
	Data   []struct {
		Mode          string `json:"mode"`
		ModeID        string `json:"mode_id"`
		Enabled       bool   `json:"enabled"`
		TeamSize      int    `json:"team_size"`
		NumberOfTeams int    `json:"number_of_teams"`
		PartySize     struct {
			Max             int   `json:"max"`
			Min             int   `json:"min"`
			Invalid         []int `json:"invalid"`
			FullPartyBypass bool  `json:"full_party_bypass"`
		} `json:"party_size"`
		HighSkill struct {
			MaxPartySize int `json:"max_party_size"`
			MinTier      int `json:"min_tier"`
			MaxTier      int `json:"max_tier"`
		} `json:"high_skill"`
		Ranked         bool `json:"ranked"`
		Tournament     bool `json:"tournament"`
		SkillDisparity []struct {
			Tier    int    `json:"tier"`
			Name    string `json:"name"`
			MaxTier struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"max_tier"`
		} `json:"skill_disparity"`
		RequiredAccountLevel int `json:"required_account_level"`
		GameRules            struct {
			OvertimeWinByTwo       bool `json:"overtime_win_by_two"`
			AllowLenientSurrender  bool `json:"allow_lenient_surrender"`
			AllowDropOut           bool `json:"allow_drop_out"`
			AssignRandomAgents     bool `json:"assign_random_agents"`
			SkipPregame            bool `json:"skip_pregame"`
			AllowOvertimeDrawVote  bool `json:"allow_overtime_draw_vote"`
			OvertimeWinByTwoCapped bool `json:"overtime_win_by_two_capped"`
			PremierMode            bool `json:"premier_mode"`
		} `json:"game_rules"`
		Platforms []string `json:"platforms"`
		Maps      []struct {
			Map struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"map"`
			Enabled bool `json:"enabled"`
		} `json:"maps"`
	} `json:"data"`
	Errors []Error `json:"errors"`
}

type GetStatusResponse struct {
	Status int    `json:"status"`
	Region string `json:"region"`
	Data   struct {
		Maintenances []struct {
			CreatedAt time.Time `json:"created_at"`
			ArchiveAt time.Time `json:"archive_at"`
			Updates   []struct {
				CreatedAt    time.Time `json:"created_at"`
				UpdatedAt    time.Time `json:"updated_at"`
				Publish      bool      `json:"publish"`
				ID           int       `json:"id"`
				Translations []struct {
					Content string `json:"content"`
					Locale  string `json:"locale"`
				} `json:"translations"`
				PublishLocations []string `json:"publish_locations"`
				Author           string   `json:"author"`
			} `json:"updates"`
			Platforms []string  `json:"platforms"`
			UpdatedAt time.Time `json:"updated_at"`
			ID        int       `json:"id"`
			Titles    []struct {
				Content string `json:"content"`
				Locale  string `json:"locale"`
			} `json:"titles"`
			MaintenanceStatus string `json:"maintenance_status"`
			IncidentSeverity  string `json:"incident_severity"`
		} `json:"maintenances"`
		Incidents []struct {
			CreatedAt time.Time `json:"created_at"`
			ArchiveAt time.Time `json:"archive_at"`
			Updates   []struct {
				CreatedAt    time.Time `json:"created_at"`
				UpdatedAt    time.Time `json:"updated_at"`
				Publish      bool      `json:"publish"`
				ID           int       `json:"id"`
				Translations []struct {
					Content string `json:"content"`
					Locale  string `json:"locale"`
				} `json:"translations"`
				PublishLocations []string `json:"publish_locations"`
				Author           string   `json:"author"`
			} `json:"updates"`
			Platforms []string  `json:"platforms"`
			UpdatedAt time.Time `json:"updated_at"`
			ID        int       `json:"id"`
			Titles    []struct {
				Content string `json:"content"`
				Locale  string `json:"locale"`
			} `json:"titles"`
			MaintenanceStatus string `json:"maintenance_status"`
			IncidentSeverity  string `json:"incident_severity"`
		} `json:"incidents"`
	} `json:"data"`
	Errors []Error `json:"errors"`
}

type GetVersionResponse struct {
	Status int `json:"status"`
	Data   struct {
		Version       string `json:"version"`
		ClientVersion string `json:"clientVersion"`
		Branch        string `json:"branch"`
		Region        string `json:"region"`
	} `json:"data"`
	Errors []Error `json:"errors"`
}

type GetRawMatchHistoryParams struct {
	Type    string `json:"type"`
	Value   string `json:"value"`
	Region  string `json:"region"`
	Queries string `json:"queries"`
}

type GetRawMatchHistoryResponse struct {
	Subject    string `json:"Subject"`
	BeginIndex int    `json:"BeginIndex"`
	EndIndex   int    `json:"EndIndex"`
	Total      int    `json:"Total"`
	History    []struct {
		MatchID       string `json:"MatchID"`
		GameStartTime int64  `json:"GameStartTime"`
		QueueID       string `json:"QueueID"`
	} `json:"History"`
	Region string  `json:"region"`
	Errors []Error `json:"errors"`
}

type GetRawMatchDetailsParams struct {
	Type    string   `json:"type"`
	Value   []string `json:"value"`
	Region  string   `json:"region"`
	Queries string   `json:"queries"`
}

type GetRawMatchDetailsResponse struct {
	MatchInfo struct {
		MatchID             string `json:"matchId"`
		MapID               string `json:"mapId"`
		GamePodID           string `json:"gamePodId"`
		GameLoopZone        string `json:"gameLoopZone"`
		GameServerAddress   string `json:"gameServerAddress"`
		GameVersion         string `json:"gameVersion"`
		GameLengthMillis    int    `json:"gameLengthMillis"`
		GameStartMillis     int64  `json:"gameStartMillis"`
		ProvisioningFlowID  string `json:"provisioningFlowID"`
		IsCompleted         bool   `json:"isCompleted"`
		CustomGameName      string `json:"customGameName"`
		ForcePostProcessing bool   `json:"forcePostProcessing"`
		QueueID             string `json:"queueID"`
		GameMode            string `json:"gameMode"`
		IsRanked            bool   `json:"isRanked"`
		IsMatchSampled      bool   `json:"isMatchSampled"`
		SeasonID            string `json:"seasonId"`
		CompletionState     string `json:"completionState"`
		PlatformType        string `json:"platformType"`
		PremierMatchInfo    struct {
		} `json:"premierMatchInfo"`
		PartyRRPenalties            map[string]int `json:"partyRRPenalties"`
		ShouldMatchDisablePenalties bool           `json:"shouldMatchDisablePenalties"`
	} `json:"matchInfo"`
	Players []struct {
		Subject      string `json:"subject"`
		GameName     string `json:"gameName"`
		TagLine      string `json:"tagLine"`
		PlatformInfo struct {
			PlatformType      string `json:"platformType"`
			PlatformOS        string `json:"platformOS"`
			PlatformOSVersion string `json:"platformOSVersion"`
			PlatformChipset   string `json:"platformChipset"`
		} `json:"platformInfo"`
		TeamID      string `json:"teamId"`
		PartyID     string `json:"partyId"`
		CharacterID string `json:"characterId"`
		Stats       struct {
			Score          int `json:"score"`
			RoundsPlayed   int `json:"roundsPlayed"`
			Kills          int `json:"kills"`
			Deaths         int `json:"deaths"`
			Assists        int `json:"assists"`
			PlaytimeMillis int `json:"playtimeMillis"`
			AbilityCasts   struct {
				GrenadeCasts  int `json:"grenadeCasts"`
				Ability1Casts int `json:"ability1Casts"`
				Ability2Casts int `json:"ability2Casts"`
				UltimateCasts int `json:"ultimateCasts"`
			} `json:"abilityCasts"`
		} `json:"stats"`
		RoundDamage []struct {
			Round    int    `json:"round"`
			Receiver string `json:"receiver"`
			Damage   int    `json:"damage"`
		} `json:"roundDamage"`
		CompetitiveTier        int    `json:"competitiveTier"`
		IsObserver             bool   `json:"isObserver"`
		PlayerCard             string `json:"playerCard"`
		PlayerTitle            string `json:"playerTitle"`
		AccountLevel           int    `json:"accountLevel"`
		SessionPlaytimeMinutes int    `json:"sessionPlaytimeMinutes"`
		BehaviorFactors        struct {
			AfkRounds                   int     `json:"afkRounds"`
			Collisions                  float64 `json:"collisions"`
			CommsRatingRecovery         int     `json:"commsRatingRecovery"`
			DamageParticipationOutgoing int     `json:"damageParticipationOutgoing"`
			FriendlyFireIncoming        int     `json:"friendlyFireIncoming"`
			FriendlyFireOutgoing        int     `json:"friendlyFireOutgoing"`
			MouseMovement               int     `json:"mouseMovement"`
			SelfDamage                  int     `json:"selfDamage"`
			StayedInSpawnRounds         int     `json:"stayedInSpawnRounds"`
		} `json:"behaviorFactors"`
		NewPlayerExperienceDetails struct {
			BasicMovement struct {
				IdleTimeMillis              int `json:"idleTimeMillis"`
				ObjectiveCompleteTimeMillis int `json:"objectiveCompleteTimeMillis"`
			} `json:"basicMovement"`
			BasicGunSkill struct {
				IdleTimeMillis              int `json:"idleTimeMillis"`
				ObjectiveCompleteTimeMillis int `json:"objectiveCompleteTimeMillis"`
			} `json:"basicGunSkill"`
			AdaptiveBots struct {
				IdleTimeMillis                               int         `json:"idleTimeMillis"`
				ObjectiveCompleteTimeMillis                  int         `json:"objectiveCompleteTimeMillis"`
				AdaptiveBotAverageDurationMillisAllAttempts  int         `json:"adaptiveBotAverageDurationMillisAllAttempts"`
				AdaptiveBotAverageDurationMillisFirstAttempt int         `json:"adaptiveBotAverageDurationMillisFirstAttempt"`
				KillDetailsFirstAttempt                      interface{} `json:"killDetailsFirstAttempt"`
			} `json:"adaptiveBots"`
			Ability struct {
				IdleTimeMillis              int `json:"idleTimeMillis"`
				ObjectiveCompleteTimeMillis int `json:"objectiveCompleteTimeMillis"`
			} `json:"ability"`
			BombPlant struct {
				IdleTimeMillis              int `json:"idleTimeMillis"`
				ObjectiveCompleteTimeMillis int `json:"objectiveCompleteTimeMillis"`
			} `json:"bombPlant"`
			DefendBombSite struct {
				IdleTimeMillis              int  `json:"idleTimeMillis"`
				ObjectiveCompleteTimeMillis int  `json:"objectiveCompleteTimeMillis"`
				Success                     bool `json:"success"`
			} `json:"defendBombSite"`
			SettingStatus struct {
				IsMouseSensitivityDefault bool `json:"isMouseSensitivityDefault"`
				IsCrosshairDefault        bool `json:"isCrosshairDefault"`
			} `json:"settingStatus"`
			VersionString string `json:"versionString"`
		} `json:"newPlayerExperienceDetails"`
		PreferredLevelBorder string `json:"preferredLevelBorder,omitempty"`
		XpModifications      []struct {
			Value float64 `json:"Value"`
			ID    string  `json:"ID"`
		} `json:"xpModifications,omitempty"`
	} `json:"players"`
	Bots    []interface{} `json:"bots"`
	Coaches []interface{} `json:"coaches"`
	Teams   []struct {
		TeamID       string `json:"teamId"`
		Won          bool   `json:"won"`
		RoundsPlayed int    `json:"roundsPlayed"`
		RoundsWon    int    `json:"roundsWon"`
		NumPoints    int    `json:"numPoints"`
	} `json:"teams"`
	RoundResults []struct {
		RoundNum             int    `json:"roundNum"`
		RoundResult          string `json:"roundResult"`
		RoundCeremony        string `json:"roundCeremony"`
		WinningTeam          string `json:"winningTeam"`
		BombPlanter          string `json:"bombPlanter,omitempty"`
		PlantRoundTime       int    `json:"plantRoundTime"`
		PlantPlayerLocations []struct {
			Subject     string  `json:"subject"`
			ViewRadians float64 `json:"viewRadians"`
			Location    struct {
				X int `json:"x"`
				Y int `json:"y"`
			} `json:"location"`
		} `json:"plantPlayerLocations"`
		PlantLocation struct {
			X int `json:"x"`
			Y int `json:"y"`
		} `json:"plantLocation"`
		PlantSite             string      `json:"plantSite"`
		DefuseRoundTime       int         `json:"defuseRoundTime"`
		DefusePlayerLocations interface{} `json:"defusePlayerLocations"`
		DefuseLocation        struct {
			X int `json:"x"`
			Y int `json:"y"`
		} `json:"defuseLocation"`
		PlayerStats []struct {
			Subject string `json:"subject"`
			Kills   []struct {
				GameTime       int    `json:"gameTime"`
				RoundTime      int    `json:"roundTime"`
				Killer         string `json:"killer"`
				Victim         string `json:"victim"`
				VictimLocation struct {
					X int `json:"x"`
					Y int `json:"y"`
				} `json:"victimLocation"`
				Assistants      []interface{} `json:"assistants"`
				PlayerLocations []struct {
					Subject     string  `json:"subject"`
					ViewRadians float64 `json:"viewRadians"`
					Location    struct {
						X int `json:"x"`
						Y int `json:"y"`
					} `json:"location"`
				} `json:"playerLocations"`
				FinishingDamage struct {
					DamageType          string `json:"damageType"`
					DamageItem          string `json:"damageItem"`
					IsSecondaryFireMode bool   `json:"isSecondaryFireMode"`
				} `json:"finishingDamage"`
			} `json:"kills"`
			Damage []struct {
				Receiver  string `json:"receiver"`
				Damage    int    `json:"damage"`
				Legshots  int    `json:"legshots"`
				Bodyshots int    `json:"bodyshots"`
				Headshots int    `json:"headshots"`
			} `json:"damage"`
			Score   int `json:"score"`
			Economy struct {
				LoadoutValue int    `json:"loadoutValue"`
				Weapon       string `json:"weapon"`
				Armor        string `json:"armor"`
				Remaining    int    `json:"remaining"`
				Spent        int    `json:"spent"`
			} `json:"economy"`
			Ability struct {
				GrenadeEffects  interface{} `json:"grenadeEffects"`
				Ability1Effects interface{} `json:"ability1Effects"`
				Ability2Effects interface{} `json:"ability2Effects"`
				UltimateEffects interface{} `json:"ultimateEffects"`
			} `json:"ability"`
			WasAfk        bool `json:"wasAfk"`
			WasPenalized  bool `json:"wasPenalized"`
			StayedInSpawn bool `json:"stayedInSpawn"`
		} `json:"playerStats"`
		RoundResultCode string `json:"roundResultCode"`
		PlayerEconomies []struct {
			Subject      string `json:"subject"`
			LoadoutValue int    `json:"loadoutValue"`
			Weapon       string `json:"weapon"`
			Armor        string `json:"armor"`
			Remaining    int    `json:"remaining"`
			Spent        int    `json:"spent"`
		} `json:"playerEconomies"`
		PlayerScores []struct {
			Subject string `json:"subject"`
			Score   int    `json:"score"`
		} `json:"playerScores"`
		BombDefuser string `json:"bombDefuser,omitempty"`
	} `json:"roundResults"`
	Kills []struct {
		GameTime       int    `json:"gameTime"`
		RoundTime      int    `json:"roundTime"`
		Round          int    `json:"round"`
		Killer         string `json:"killer"`
		Victim         string `json:"victim"`
		VictimLocation struct {
			X int `json:"x"`
			Y int `json:"y"`
		} `json:"victimLocation"`
		Assistants      []interface{} `json:"assistants"`
		PlayerLocations []struct {
			Subject     string  `json:"subject"`
			ViewRadians float64 `json:"viewRadians"`
			Location    struct {
				X int `json:"x"`
				Y int `json:"y"`
			} `json:"location"`
		} `json:"playerLocations"`
		FinishingDamage struct {
			DamageType          string `json:"damageType"`
			DamageItem          string `json:"damageItem"`
			IsSecondaryFireMode bool   `json:"isSecondaryFireMode"`
		} `json:"finishingDamage"`
	} `json:"kills"`
	Region string  `json:"region"`
	Error bool `json:"error"`
	Code int `json:"code"`
	Id string `json:"id"`
}

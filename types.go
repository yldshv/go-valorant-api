package govapi

import "time"

type GetAccountByNameParams struct {
	Name string
	Tag  string
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
	Puuid string
	Mode string
	Map string
	Size string
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
	Puuid string
	Season string
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
				Error            string   `json:"error"`
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
				Error            string   `json:"error"`
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
				Error            string   `json:"error"`
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
				Error            string   `json:"error"`
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
				Error            string   `json:"error"`
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
				Error            string   `json:"error"`
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
				Error            string   `json:"error"`
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
				Error            string   `json:"error"`
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
				Error            string   `json:"error"`
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
				Error            string   `json:"error"`
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
				Error            string   `json:"error"`
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
				Error            string   `json:"error"`
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
				Error            string   `json:"error"`
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
				Error            string   `json:"error"`
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
				Error            string   `json:"error"`
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
				Error            string   `json:"error"`
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
				Error            string   `json:"error"`
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
				Error            string   `json:"error"`
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
	Puuid string
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
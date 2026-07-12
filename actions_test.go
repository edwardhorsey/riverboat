//* Copyright (c) 2020, Alex Lewontin
//* All rights reserved.
//*
//* Redistribution and use in source and binary forms, with or without
//* modification, are permitted provided that the following conditions are met:
//*
//* - Redistributions of source code must retain the above copyright notice, this
//* list of conditions and the following disclaimer.
//* - Redistributions in binary form must reproduce the above copyright notice,
//* this list of conditions and the following disclaimer in the documentation
//* and/or other materials provided with the distribution.
//*
//* THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
//* ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
//* WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
//* DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
//* FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
//* DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
//* SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
//* CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
//* OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
//* OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package riverboat

import (
	"testing"
)

func TestIntegration_Scenarios(t *testing.T) {

	t.Run("Scenario 1 deal none ready", func(t *testing.T) {
		g := NewGame()

		pn_a := g.AddPlayer()
		g.AddPlayer()
		g.AddPlayer()

		err := Deal(g, pn_a, 0)

		if err != ErrIllegalAction {
			t.Error("Test failed - Deal must return ErrIllegalAction as 0 players are marked ready.")
		}
	})

	t.Run("Scenario 2 ready", func(t *testing.T) {
		g := NewGame()

		pn_a := g.AddPlayer()
		g.AddPlayer()
		g.AddPlayer()

		err := ToggleReady(g, pn_a, 0)

		if err != ErrIllegalAction {
			t.Error("Test failed - ToggleReady must return ErrIllegalAction as player 0 has not bought in.")
		}
	})

	t.Run("Scenario 3 deal", func(t *testing.T) {
		var err error
		g := NewGame()

		pn_a := g.AddPlayer()
		pn_b := g.AddPlayer()
		pn_c := g.AddPlayer()

		err = BuyIn(g, pn_a, 100)

		if err != nil {
			t.Errorf("Test failed - Error buying in: %s", err)
		}

		err = BuyIn(g, pn_b, 100)

		if err != nil {
			t.Errorf("Test failed - Error buying in: %s", err)
		}

		err = BuyIn(g, pn_c, 100)

		if err != nil {
			t.Errorf("Test failed - Error buying in: %s", err)
		}

		err = ToggleReady(g, pn_a, 0)

		if err != nil {
			t.Errorf("Test failed - Error marking ready: %s", err)
		}

		err = ToggleReady(g, pn_b, 0)

		if err != nil {
			t.Errorf("Test failed - Error marking ready: %s", err)
		}

		err = ToggleReady(g, pn_c, 0)

		if err != nil {
			t.Errorf("Test failed - Error marking ready: %s", err)
		}

		err = Deal(g, pn_a, 0)

		if err != nil {
			t.Errorf("Test failed - error dealing: %s", err)
		}
	})

	t.Run("Scenario 4 deal check for wrong dealer", func(t *testing.T) {
		var err error
		g := NewGame()

		pn_a := g.AddPlayer()
		pn_b := g.AddPlayer()
		pn_c := g.AddPlayer()

		err = BuyIn(g, pn_a, 100)

		if err != nil {
			t.Errorf("Test failed - Error buying in: %s", err)
		}

		err = BuyIn(g, pn_b, 100)

		if err != nil {
			t.Errorf("Test failed - Error buying in: %s", err)
		}

		err = BuyIn(g, pn_c, 100)

		if err != nil {
			t.Errorf("Test failed - Error buying in: %s", err)
		}

		err = ToggleReady(g, pn_a, 0)

		if err != nil {
			t.Errorf("Test failed - Error marking ready: %s", err)
		}

		err = ToggleReady(g, pn_b, 0)

		if err != nil {
			t.Errorf("Test failed - Error marking ready: %s", err)
		}

		err = ToggleReady(g, pn_c, 0)

		if err != nil {
			t.Errorf("Test failed - Error marking ready: %s", err)
		}

		err = Deal(g, pn_b, 0)

		if err != ErrIllegalAction {
			t.Errorf("Test failed - must return ErrIllegalAction as pn_b is not the dealer")
		}
	})

	t.Run("Scenario 5 bet success", func(t *testing.T) {
		var err error
		g := NewGame()

		pn_a := g.AddPlayer()
		pn_b := g.AddPlayer()
		pn_c := g.AddPlayer()

		err = BuyIn(g, pn_a, 100)

		if err != nil {
			t.Errorf("Test failed - Error buying in: %s", err)
		}

		err = BuyIn(g, pn_b, 100)

		if err != nil {
			t.Errorf("Test failed - Error buying in: %s", err)
		}

		err = BuyIn(g, pn_c, 100)

		if err != nil {
			t.Errorf("Test failed - Error buying in: %s", err)
		}

		err = ToggleReady(g, pn_a, 0)

		if err != nil {
			t.Errorf("Test failed - Error marking ready: %s", err)
		}

		err = ToggleReady(g, pn_b, 0)

		if err != nil {
			t.Errorf("Test failed - Error marking ready: %s", err)
		}

		err = ToggleReady(g, pn_c, 0)

		if err != nil {
			t.Errorf("Test failed - Error marking ready: %s", err)
		}

		err = Deal(g, pn_a, 0)

		if err != nil {
			t.Errorf("Test failed - error dealing: %s", err)
		}

		err = Bet(g, pn_a, 25)

		if err != nil {

			t.Errorf("Test failed - error betting: %s", err)
		}

		if g.players[pn_a].Bet != 25 {
			t.Errorf("Betting mechanic not working.")
		}
	})

	t.Run("Scenario 6 simple", func(t *testing.T) {
		var err error
		g := NewGame()

		pn_a := g.AddPlayer()
		pn_b := g.AddPlayer()
		pn_c := g.AddPlayer()

		err = BuyIn(g, pn_a, 100)

		if err != nil {
			t.Errorf("Test failed - Error buying in: %s", err)
		}

		err = BuyIn(g, pn_b, 100)

		if err != nil {
			t.Errorf("Test failed - Error buying in: %s", err)
		}

		err = BuyIn(g, pn_c, 100)

		if err != nil {
			t.Errorf("Test failed - Error buying in: %s", err)
		}

		err = ToggleReady(g, pn_a, 0)

		if err != nil {
			t.Errorf("Test failed - Error marking ready: %s", err)
		}

		err = ToggleReady(g, pn_b, 0)

		if err != nil {
			t.Errorf("Test failed - Error marking ready: %s", err)
		}

		err = ToggleReady(g, pn_c, 0)

		if err != nil {
			t.Errorf("Test failed - Error marking ready: %s", err)
		}

		// Preflop

		err = Deal(g, pn_a, 0)

		if err != nil {
			t.Errorf("Test failed - error dealing: %s", err)
		}

		err = Bet(g, pn_a, 25)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		err = Bet(g, pn_b, 15)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		err = Bet(g, pn_c, 0)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		// Flop

		err = Bet(g, pn_b, 25)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		err = Bet(g, pn_c, 25)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		err = Bet(g, pn_a, 25)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		// Turn

		err = Bet(g, pn_b, 0)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		err = Bet(g, pn_c, 0)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		err = Bet(g, pn_a, 0)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		//River
		err = Bet(g, pn_b, 0)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		err = Bet(g, pn_c, 0)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		err = Bet(g, pn_a, 0)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

	})

	t.Run("Scenario 7 fold", func(t *testing.T) {
		var err error
		g := NewGame()

		pn_a := g.AddPlayer()
		pn_b := g.AddPlayer()
		pn_c := g.AddPlayer()

		err = BuyIn(g, pn_a, 100)

		if err != nil {
			t.Errorf("Test failed - Error buying in: %s", err)
		}

		err = BuyIn(g, pn_b, 100)

		if err != nil {
			t.Errorf("Test failed - Error buying in: %s", err)
		}

		err = BuyIn(g, pn_c, 100)

		if err != nil {
			t.Errorf("Test failed - Error buying in: %s", err)
		}

		err = ToggleReady(g, pn_a, 0)

		if err != nil {
			t.Errorf("Test failed - Error marking ready: %s", err)
		}

		err = ToggleReady(g, pn_b, 0)

		if err != nil {
			t.Errorf("Test failed - Error marking ready: %s", err)
		}

		err = ToggleReady(g, pn_c, 0)

		if err != nil {
			t.Errorf("Test failed - Error marking ready: %s", err)
		}

		// Preflop

		err = Deal(g, pn_a, 0)

		if err != nil {
			t.Errorf("Test failed - error dealing: %s", err)
		}

		err = Bet(g, pn_a, 25)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		err = Bet(g, pn_b, 15)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		err = Bet(g, pn_c, 0)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		// Flop

		err = Bet(g, pn_b, 25)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		err = Fold(g, pn_c, 0)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		err = Bet(g, pn_a, 25)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		// Turn
		err = Bet(g, pn_b, 0)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		err = Bet(g, pn_a, 0)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		//River

		err = Bet(g, pn_b, 0)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		err = Bet(g, pn_a, 0)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

	})

	t.Run("Scenario 8 reraise", func(t *testing.T) {
		var err error
		g := NewGame()

		pn_a := g.AddPlayer()
		pn_b := g.AddPlayer()
		pn_c := g.AddPlayer()

		err = BuyIn(g, pn_a, 100)

		if err != nil {
			t.Errorf("Test failed - Error buying in: %s", err)
		}

		err = BuyIn(g, pn_b, 100)

		if err != nil {
			t.Errorf("Test failed - Error buying in: %s", err)
		}

		err = BuyIn(g, pn_c, 100)

		if err != nil {
			t.Errorf("Test failed - Error buying in: %s", err)
		}

		err = ToggleReady(g, pn_a, 0)

		if err != nil {
			t.Errorf("Test failed - Error marking ready: %s", err)
		}

		err = ToggleReady(g, pn_b, 0)

		if err != nil {
			t.Errorf("Test failed - Error marking ready: %s", err)
		}

		err = ToggleReady(g, pn_c, 0)

		if err != nil {
			t.Errorf("Test failed - Error marking ready: %s", err)
		}

		// Preflop

		err = Deal(g, pn_a, 0)

		if err != nil {
			t.Errorf("Test failed - error dealing: %s", err)
		}

		err = Bet(g, pn_a, 25)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		err = Bet(g, pn_b, 15)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		err = Bet(g, pn_c, 0)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		// Flop

		err = Bet(g, pn_b, 25)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		err = Fold(g, pn_c, 0)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		err = Bet(g, pn_a, 50)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		err = Bet(g, pn_b, 25)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		// Turn

		err = Bet(g, pn_b, 0)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		err = Bet(g, pn_a, 0)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		//River

		err = Bet(g, pn_b, 0)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		err = Bet(g, pn_a, 0)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

	})

	t.Run("Scenario 9 leave", func(t *testing.T) {
		g := NewGame()

		pnA := g.AddPlayer()
		pnB := g.AddPlayer()

		err := BuyIn(g, pnA, 100)
		if err != nil {
			t.Fatalf("BuyIn pnA: %v", err)
		}
		err = BuyIn(g, pnB, 100)
		if err != nil {
			t.Fatalf("BuyIn pnB: %v", err)
		}

		err = ToggleReady(g, pnA, 0)
		if err != nil {
			t.Fatalf("ToggleReady pnA: %v", err)
		}

		view := g.GenerateOmniView()
		if !view.Players[pnA].Ready {
			t.Fatal("expected pnA ready before leave")
		}
		if view.Players[pnA].Left {
			t.Fatal("expected pnA not left before leave")
		}

		err = Leave(g, pnA, 0)
		if err != nil {
			t.Fatalf("Leave pnA: %v", err)
		}

		view = g.GenerateOmniView()
		if view.Players[pnA].Ready {
			t.Error("expected pnA not ready after leave")
		}
		if !view.Players[pnA].Left {
			t.Error("expected pnA left after leave")
		}
	})

	t.Run("Scenario 10 ready count", func(t *testing.T) {
		g := NewGame()

		pnA := g.AddPlayer()
		pnB := g.AddPlayer()
		pnC := g.AddPlayer()

		err := BuyIn(g, pnA, 100)
		if err != nil {
			t.Fatalf("BuyIn pnA: %v", err)
		}
		err = BuyIn(g, pnB, 100)
		if err != nil {
			t.Fatalf("BuyIn pnB: %v", err)
		}
		err = BuyIn(g, pnC, 100)
		if err != nil {
			t.Fatalf("BuyIn pnC: %v", err)
		}

		err = ToggleReady(g, pnA, 0)
		if err != nil {
			t.Fatalf("ToggleReady pnA: %v", err)
		}
		err = ToggleReady(g, pnB, 0)
		if err != nil {
			t.Fatalf("ToggleReady pnB: %v", err)
		}
		err = ToggleReady(g, pnC, 0)
		if err != nil {
			t.Fatalf("ToggleReady pnC: %v", err)
		}

		if g.readyCount() != 3 {
			t.Errorf("expected readyCount 3, got %d", g.readyCount())
		}

		err = Leave(g, pnB, 0)
		if err != nil {
			t.Fatalf("Leave pnB: %v", err)
		}

		if g.readyCount() != 2 {
			t.Errorf("expected readyCount 2 after leave, got %d", g.readyCount())
		}
	})

	t.Run("Scenario 11 reuse seat", func(t *testing.T) {
		g := NewGame()

		pnA := g.AddPlayer()
		pnB := g.AddPlayer()
		pnC := g.AddPlayer()

		err := BuyIn(g, pnA, 100)
		if err != nil {
			t.Fatalf("BuyIn pnA: %v", err)
		}
		err = BuyIn(g, pnB, 100)
		if err != nil {
			t.Fatalf("BuyIn pnB: %v", err)
		}
		err = BuyIn(g, pnC, 100)
		if err != nil {
			t.Fatalf("BuyIn pnC: %v", err)
		}

		err = Leave(g, pnB, 0)
		if err != nil {
			t.Fatalf("Leave pnB: %v", err)
		}

		if len(g.players) != 3 {
			t.Fatalf("expected 3 players before reuse, got %d", len(g.players))
		}

		pnD := g.AddPlayer()

		if len(g.players) != 3 {
			t.Errorf("expected 3 players after reuse, got %d", len(g.players))
		}
		if pnD != pnB {
			t.Errorf("expected new player to reuse pnB slot (%d), got %d", pnB, pnD)
		}

		view := g.GenerateOmniView()
		if view.Players[pnD].Left {
			t.Error("expected reused player to not be left")
		}
		if view.Players[pnD].Ready {
			t.Error("expected reused player to not be ready")
		}
		if view.Players[pnD].Stack != 0 {
			t.Errorf("expected reused player stack 0, got %d", view.Players[pnD].Stack)
		}
	})

	t.Run("Scenario 12 deal with left players", func(t *testing.T) {
		g := NewGame()

		pnA := g.AddPlayer()
		pnB := g.AddPlayer()
		pnC := g.AddPlayer()

		err := BuyIn(g, pnA, 100)
		if err != nil {
			t.Fatalf("BuyIn pnA: %v", err)
		}
		err = BuyIn(g, pnB, 100)
		if err != nil {
			t.Fatalf("BuyIn pnB: %v", err)
		}
		err = BuyIn(g, pnC, 100)
		if err != nil {
			t.Fatalf("BuyIn pnC: %v", err)
		}

		err = ToggleReady(g, pnA, 0)
		if err != nil {
			t.Fatalf("ToggleReady pnA: %v", err)
		}
		err = ToggleReady(g, pnB, 0)
		if err != nil {
			t.Fatalf("ToggleReady pnB: %v", err)
		}
		err = ToggleReady(g, pnC, 0)
		if err != nil {
			t.Fatalf("ToggleReady pnC: %v", err)
		}

		err = Leave(g, pnB, 0)
		if err != nil {
			t.Fatalf("Leave pnB: %v", err)
		}

		err = Deal(g, pnA, 0)
		if err != nil {
			t.Fatalf("Deal: %v", err)
		}

		view := g.GenerateOmniView()
		if !view.Players[pnA].In {
			t.Error("expected pnA to be in after deal")
		}
		if view.Players[pnA].Cards[0] == 0 || view.Players[pnA].Cards[1] == 0 {
			t.Error("expected pnA to have cards")
		}
		if view.Players[pnB].In {
			t.Error("expected pnB (left) to not be in after deal")
		}
		if view.Players[pnB].Cards[0] != 0 || view.Players[pnB].Cards[1] != 0 {
			t.Error("expected pnB (left) to have no cards")
		}
		if !view.Players[pnC].In {
			t.Error("expected pnC to be in after deal")
		}
		if view.Players[pnC].Cards[0] == 0 || view.Players[pnC].Cards[1] == 0 {
			t.Error("expected pnC to have cards")
		}
	})

	t.Run("Scenario 13 blinds with left players", func(t *testing.T) {
		g := NewGame()

		pnA := g.AddPlayer()
		pnB := g.AddPlayer()
		pnC := g.AddPlayer()

		err := BuyIn(g, pnA, 100)
		if err != nil {
			t.Fatalf("BuyIn pnA: %v", err)
		}
		err = BuyIn(g, pnB, 100)
		if err != nil {
			t.Fatalf("BuyIn pnB: %v", err)
		}
		err = BuyIn(g, pnC, 100)
		if err != nil {
			t.Fatalf("BuyIn pnC: %v", err)
		}

		err = ToggleReady(g, pnA, 0)
		if err != nil {
			t.Fatalf("ToggleReady pnA: %v", err)
		}
		err = ToggleReady(g, pnB, 0)
		if err != nil {
			t.Fatalf("ToggleReady pnB: %v", err)
		}
		err = ToggleReady(g, pnC, 0)
		if err != nil {
			t.Fatalf("ToggleReady pnC: %v", err)
		}

		err = Leave(g, pnB, 0)
		if err != nil {
			t.Fatalf("Leave pnB: %v", err)
		}

		g.updateBlindNums()

		if g.sbNum == pnB {
			t.Error("expected SB to not be left player pnB")
		}
		if g.bbNum == pnB {
			t.Error("expected BB to not be left player pnB")
		}
		if g.utgNum == pnB {
			t.Error("expected UTG to not be left player pnB")
		}
	})

	t.Run("Scenario 14 dealer skip with left players", func(t *testing.T) {
		g := NewGame()

		pnA := g.AddPlayer()
		pnB := g.AddPlayer()

		err := BuyIn(g, pnA, 100)
		if err != nil {
			t.Fatalf("BuyIn pnA: %v", err)
		}
		err = BuyIn(g, pnB, 100)
		if err != nil {
			t.Fatalf("BuyIn pnB: %v", err)
		}

		err = ToggleReady(g, pnA, 0)
		if err != nil {
			t.Fatalf("ToggleReady pnA: %v", err)
		}
		err = ToggleReady(g, pnB, 0)
		if err != nil {
			t.Fatalf("ToggleReady pnB: %v", err)
		}

		err = Deal(g, pnA, 0)
		if err != nil {
			t.Fatalf("Deal: %v", err)
		}

		view := g.GenerateOmniView()
		if view.Stage != 2 {
			t.Fatalf("expected PreFlop, got %d", view.Stage)
		}

		actionNum := view.ActionNum
		err = Fold(g, actionNum, 0)
		if err != nil {
			t.Fatalf("Fold actionNum %d: %v", actionNum, err)
		}

		view = g.GenerateOmniView()
		if view.Stage != 1 {
			t.Fatalf("expected PreDeal after fold, got %d", view.Stage)
		}

		err = Leave(g, pnA, 0)
		if err != nil {
			t.Fatalf("Leave pnA: %v", err)
		}

		g.resetForNextHand()

		if g.dealerNum == pnA {
			t.Error("expected dealer to not be left player pnA")
		}
		if !g.players[pnB].Ready {
			t.Error("expected pnB to still be ready after reset")
		}
		if g.players[pnA].In {
			t.Error("expected left player pnA to not be in after reset")
		}
	})

	t.Run("Scenario 15 rejoin", func(t *testing.T) {
		g := NewGame()

		pnA := g.AddPlayer()
		pnB := g.AddPlayer()
		pnC := g.AddPlayer()

		err := BuyIn(g, pnA, 100)
		if err != nil {
			t.Fatalf("BuyIn pnA: %v", err)
		}
		err = BuyIn(g, pnB, 100)
		if err != nil {
			t.Fatalf("BuyIn pnB: %v", err)
		}
		err = BuyIn(g, pnC, 100)
		if err != nil {
			t.Fatalf("BuyIn pnC: %v", err)
		}

		err = ToggleReady(g, pnA, 0)
		if err != nil {
			t.Fatalf("ToggleReady pnA: %v", err)
		}
		err = ToggleReady(g, pnB, 0)
		if err != nil {
			t.Fatalf("ToggleReady pnB: %v", err)
		}
		err = ToggleReady(g, pnC, 0)
		if err != nil {
			t.Fatalf("ToggleReady pnC: %v", err)
		}

		err = Deal(g, pnA, 0)
		if err != nil {
			t.Fatalf("Deal hand 1: %v", err)
		}

		err = Bet(g, pnA, 25)
		if err != nil {
			t.Fatalf("Bet pnA: %v", err)
		}
		err = Bet(g, pnB, 15)
		if err != nil {
			t.Fatalf("Bet pnB: %v", err)
		}
		err = Bet(g, pnC, 0)
		if err != nil {
			t.Fatalf("Bet pnC: %v", err)
		}

		err = Bet(g, pnB, 25)
		if err != nil {
			t.Fatalf("Bet pnB flop: %v", err)
		}
		err = Bet(g, pnC, 25)
		if err != nil {
			t.Fatalf("Bet pnC flop: %v", err)
		}
		err = Bet(g, pnA, 25)
		if err != nil {
			t.Fatalf("Bet pnA flop: %v", err)
		}

		err = Bet(g, pnB, 0)
		if err != nil {
			t.Fatalf("Bet pnB turn: %v", err)
		}
		err = Bet(g, pnC, 0)
		if err != nil {
			t.Fatalf("Bet pnC turn: %v", err)
		}
		err = Bet(g, pnA, 0)
		if err != nil {
			t.Fatalf("Bet pnA turn: %v", err)
		}

		err = Bet(g, pnB, 0)
		if err != nil {
			t.Fatalf("Bet pnB river: %v", err)
		}
		err = Bet(g, pnC, 0)
		if err != nil {
			t.Fatalf("Bet pnC river: %v", err)
		}
		err = Bet(g, pnA, 0)
		if err != nil {
			t.Fatalf("Bet pnA river: %v", err)
		}

		view := g.GenerateOmniView()
		if view.Stage != 1 {
			t.Fatalf("expected stage PreDeal after hand, got %d", view.Stage)
		}

		err = Leave(g, pnB, 0)
		if err != nil {
			t.Fatalf("Leave pnB: %v", err)
		}

		pnD := g.AddPlayer()
		if pnD != pnB {
			t.Errorf("expected new player to reuse pnB slot, got %d", pnD)
		}

		err = BuyIn(g, pnD, 100)
		if err != nil {
			t.Fatalf("BuyIn pnD: %v", err)
		}

		err = ToggleReady(g, pnD, 0)
		if err != nil {
			t.Fatalf("ToggleReady pnD: %v", err)
		}

		if g.readyCount() != 3 {
			t.Errorf("expected readyCount 3 (pnA, pnC, pnD), got %d", g.readyCount())
		}

		err = Deal(g, g.dealerNum, 0)
		if err != nil {
			t.Fatalf("Deal hand 2: %v", err)
		}

		view = g.GenerateOmniView()
		if view.Stage != 2 {
			t.Errorf("expected stage PreFlop after deal, got %d", view.Stage)
		}
		if !view.Betting {
			t.Error("expected Betting=true after deal")
		}
	})

	t.Run("Scenario 16 blocked actions for left players", func(t *testing.T) {
		g := NewGame()

		pnA := g.AddPlayer()
		pnB := g.AddPlayer()
		pnC := g.AddPlayer()

		err := BuyIn(g, pnA, 100)
		if err != nil {
			t.Fatalf("BuyIn pnA: %v", err)
		}
		err = BuyIn(g, pnB, 100)
		if err != nil {
			t.Fatalf("BuyIn pnB: %v", err)
		}
		err = BuyIn(g, pnC, 100)
		if err != nil {
			t.Fatalf("BuyIn pnC: %v", err)
		}

		err = ToggleReady(g, pnA, 0)
		if err != nil {
			t.Fatalf("ToggleReady pnA: %v", err)
		}
		err = ToggleReady(g, pnB, 0)
		if err != nil {
			t.Fatalf("ToggleReady pnB: %v", err)
		}
		err = ToggleReady(g, pnC, 0)
		if err != nil {
			t.Fatalf("ToggleReady pnC: %v", err)
		}

		err = Leave(g, pnA, 0)
		if err != nil {
			t.Fatalf("Leave pnA: %v", err)
		}

		err = Deal(g, pnB, 0)
		if err != nil {
			t.Fatalf("Deal: %v", err)
		}

		err = Bet(g, pnA, 25)
		if err == nil {
			t.Error("expected Bet to fail for left player")
		}

		err = Fold(g, pnA, 0)
		if err == nil {
			t.Error("expected Fold to fail for left player")
		}

		view := g.GenerateOmniView()
		if view.Players[pnA].Left {
			t.Log("pnA is left, ToggleReady should work (coming back)")
		}

		err = ToggleReady(g, pnA, 0)
		if err != nil {
			t.Errorf("expected ToggleReady to succeed for left player with stack, got: %v", err)
		}

		view = g.GenerateOmniView()
		if view.Players[pnA].Left {
			t.Error("expected pnA to no longer be left after ToggleReady")
		}
		if !view.Players[pnA].Ready {
			t.Error("expected pnA to be ready after ToggleReady")
		}
	})

	t.Run("Scenario 17 fold concedes winner", func(t *testing.T) {
		g := NewGame()

		pnA := g.AddPlayer()
		pnB := g.AddPlayer()

		err := BuyIn(g, pnA, 100)
		if err != nil {
			t.Fatalf("BuyIn pnA: %v", err)
		}
		err = BuyIn(g, pnB, 100)
		if err != nil {
			t.Fatalf("BuyIn pnB: %v", err)
		}

		err = ToggleReady(g, pnA, 0)
		if err != nil {
			t.Fatalf("ToggleReady pnA: %v", err)
		}
		err = ToggleReady(g, pnB, 0)
		if err != nil {
			t.Fatalf("ToggleReady pnB: %v", err)
		}

		stackA := g.players[pnA].Stack
		stackB := g.players[pnB].Stack

		err = Deal(g, pnA, 0)
		if err != nil {
			t.Fatalf("Deal: %v", err)
		}

		actionNum := g.GenerateOmniView().ActionNum
		err = Fold(g, actionNum, 0)
		if err != nil {
			t.Fatalf("Fold: %v", err)
		}

		view := g.GenerateOmniView()

		if view.Stage != PreDeal {
			t.Errorf("expected PreDeal after fold concession, got %d", view.Stage)
		}

		if len(view.Pots) == 0 {
			t.Fatal("expected pots to have winner info after fold concession")
		}

		winner := (actionNum + 1) % 2
		if len(view.Pots[0].WinningPlayerNums) != 1 {
			t.Fatalf("expected 1 winner in pot, got %d", len(view.Pots[0].WinningPlayerNums))
		}
		if view.Pots[0].WinningPlayerNums[0] != winner {
			t.Errorf("expected winner to be player %d, got %d", winner, view.Pots[0].WinningPlayerNums[0])
		}

		if view.Pots[0].Amt == 0 {
			t.Error("expected pot to have money")
		}

		if winner == 0 {
			if g.players[pnA].Stack <= stackA {
				t.Errorf("expected winner (pnA) stack to increase, was %d, now %d", stackA, g.players[pnA].Stack)
			}
		} else {
			if g.players[pnB].Stack <= stackB {
				t.Errorf("expected winner (pnB) stack to increase, was %d, now %d", stackB, g.players[pnB].Stack)
			}
		}
	})

	t.Run("Scenario 18 fold concedes winner with side pots", func(t *testing.T) {
		g := NewGame()

		pnA := g.AddPlayer()
		pnB := g.AddPlayer()
		pnC := g.AddPlayer()

		err := BuyIn(g, pnA, 1000)
		if err != nil {
			t.Fatalf("BuyIn pnA: %v", err)
		}
		err = BuyIn(g, pnB, 1000)
		if err != nil {
			t.Fatalf("BuyIn pnB: %v", err)
		}
		err = BuyIn(g, pnC, 1000)
		if err != nil {
			t.Fatalf("BuyIn pnC: %v", err)
		}

		err = ToggleReady(g, pnA, 0)
		if err != nil {
			t.Fatalf("ToggleReady pnA: %v", err)
		}
		err = ToggleReady(g, pnB, 0)
		if err != nil {
			t.Fatalf("ToggleReady pnB: %v", err)
		}
		err = ToggleReady(g, pnC, 0)
		if err != nil {
			t.Fatalf("ToggleReady pnC: %v", err)
		}

		err = Deal(g, pnA, 0)
		if err != nil {
			t.Fatalf("Deal: %v", err)
		}

		actionNum := g.GenerateOmniView().ActionNum

		err = Fold(g, actionNum, 0)
		if err != nil {
			t.Fatalf("Fold pnA: %v", err)
		}

		view := g.GenerateOmniView()
		actionNum = view.ActionNum
		err = Fold(g, actionNum, 0)
		if err != nil {
			t.Fatalf("Fold pnB: %v", err)
		}

		view = g.GenerateOmniView()

		if view.Stage != PreDeal {
			t.Errorf("expected PreDeal after fold concession, got %d", view.Stage)
		}

		if len(view.Pots) == 0 {
			t.Fatal("expected pots to have winner info after fold concession")
		}

		for i, pot := range view.Pots {
			if len(pot.WinningPlayerNums) != 1 {
				t.Fatalf("pot %d: expected 1 winner, got %d", i, len(pot.WinningPlayerNums))
			}
			if pot.WinningPlayerNums[0] != pnC {
				t.Errorf("pot %d: expected winner to be pnC (%d), got %d", i, pnC, pot.WinningPlayerNums[0])
			}
		}
	})

	t.Run("Scenario 19 preserve cards through showdown", func(t *testing.T) {
		g := NewGame()

		pnA := g.AddPlayer()
		pnB := g.AddPlayer()

		for _, playerNum := range []uint{pnA, pnB} {
			if err := BuyIn(g, playerNum, 100); err != nil {
				t.Fatalf("BuyIn player %d: %v", playerNum, err)
			}
			if err := ToggleReady(g, playerNum, 0); err != nil {
				t.Fatalf("ToggleReady player %d: %v", playerNum, err)
			}
		}

		if err := Deal(g, pnA, 0); err != nil {
			t.Fatalf("Deal: %v", err)
		}

		dealtCardsA := g.players[pnA].Cards
		dealtCardsB := g.players[pnB].Cards

		for g.getStage() != PreDeal {
			actionNum := g.actionNum
			callAmount := g.toCall() - g.players[actionNum].Bet
			if err := Bet(g, actionNum, callAmount); err != nil {
				t.Fatalf("Bet player %d: %v", actionNum, err)
			}
		}

		view := g.GenerateOmniView()
		if view.Players[pnA].Cards != dealtCardsA {
			t.Errorf("expected player %d cards to remain after showdown", pnA)
		}
		if view.Players[pnB].Cards != dealtCardsB {
			t.Errorf("expected player %d cards to remain after showdown", pnB)
		}

		if err := Deal(g, g.dealerNum, 0); err != nil {
			t.Fatalf("Deal next hand: %v", err)
		}

		view = g.GenerateOmniView()
		if view.Players[pnA].Cards == dealtCardsA {
			t.Errorf("expected player %d cards to be replaced on next deal", pnA)
		}
		if view.Players[pnB].Cards == dealtCardsB {
			t.Errorf("expected player %d cards to be replaced on next deal", pnB)
		}
	})

}

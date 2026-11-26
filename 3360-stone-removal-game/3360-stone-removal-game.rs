impl Solution {
    // Time: O(1) - at most 10 iterations | Space: O(1)
    pub fn can_alice_win(n: i32) -> bool {
        let (mut stones, mut take) = (n, 10);
        let mut alice_turn = true;

        // Current player must remove exactly `take` stones; stop if impossible
        while take > 0 && stones >= take {
            stones -= take;
            take -= 1;
            alice_turn = !alice_turn;
        }

        // The player whose turn it is cannot move and loses.
        // Alice wins iff it is Bob's turn at this point.
        !alice_turn
    }
}
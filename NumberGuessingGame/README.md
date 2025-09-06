---

📌 Features

Random number generation => math/rand

User input handling => fmt.Scan

Concurrency with goroutines => separate input listener

Timeout handling with channels and time.After

Clean CLI-based interface



---

How to Play

The computer picks a random number between 1 and 50.
You have 15 seconds to guess the correct number.
After each guess:

If your guess is lower than the chosen number => "Too low!" → guess higher.

If your guess is lower than the chosen number =>  "Too high!" → guess lower.
If you guess correctly → You win!
If time runs out → You lose.
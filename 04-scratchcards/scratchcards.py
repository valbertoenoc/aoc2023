import sys
from collections import defaultdict

with open("input.txt", "r") as f:
    data = f.read()

lines = data.split("\n")
lines.pop()  # remove last element, which is an ampty line

card_value = []
N = defaultdict(int)
for i, l in enumerate(lines):
    N[i] += 1
    games = l.split("|")
    winning_numbers = games[0].split(":")[1].split()
    my_numbers = games[1].split()

    total = winning_numbers + my_numbers
    set_total = set(total)
    n_matches = len(total) - len(set_total)

    # print(f"Winning games: {winning_games}")
    # print(f"My numbers: {my_numbers}")

    if n_matches > 0:
        card_value.append(2 ** (n_matches - 1))

    for j in range(n_matches):
        N[i + 1 + j] += N[i]

print(card_value)
print(sum(card_value))
print(N)
print(sum(N.values()))

import sys

with open("input.txt", "r") as f:
    data = f.read()

lines = data.split("\n")
lines.pop()  # remove last element, which is an ampty line

card_value = []
for l in lines:
    games = l.split("|")
    winning_numbers = games[0].split(":")[1]
    my_numbers = games[1]

    # print(f"Winning games: {winning_games}")
    # print(f"My numbers: {my_numbers}")

    sum_matches = 0
    for w in winning_numbers.split():
        print(winning_numbers)
        print(w)
        print(my_numbers)
        if w in my_numbers:
            sum_matches += 1
    print("*" * 20)
    print(sum_matches)
    sum_after_first = 0
    for i in range(sum_matches):
        if i == 0:
            sum_after_first += 1
            continue

        sum_after_first *= 2

    card_value.append(sum_after_first)

print(card_value)
print(sum(card_value))

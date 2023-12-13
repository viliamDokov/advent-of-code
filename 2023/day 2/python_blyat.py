import re

numbers = "1234567890"
sumOfGames = 0
with open("inp") as f:
    for line in f:
        game = line.split(":")

        gameID = re.search("\d+", game[0])
        # print(gameID.group())
        gameValid = True
        game_numbers = game[1].split(";")
        mostRed = 0
        mostBlue = 0
        mostGreen = 0
        for subgame in game_numbers:
            pattern = re.compile(r"(\d+)(.{2})")
            result = re.findall(pattern, subgame)
            for balls in result:
                ballsAmount = int(balls[0])
                # print(balls[1])
                # print(ballsAmount)
                if balls[1] == " r" and ballsAmount > mostRed:
                    print(1)
                    mostRed = ballsAmount
                if balls[1] == " b" and ballsAmount > mostBlue:
                    print(2)
                    mostBlue = ballsAmount
                if balls[1] == " g" and ballsAmount > mostGreen:
                    print(3)
                    mostGreen = ballsAmount

        sumOfGames += mostBlue * mostGreen * mostRed
        # print(result)
        # for number in result:
        #     print(number.start())
        # for character in game:
        #     if (character in numbers):

print(sumOfGames)

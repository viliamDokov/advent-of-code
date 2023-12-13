def main():
    with open("inp") as f:
        s = f.read()

        for i in range(14, len(s)):
            window = set(s[i-14:i])
            print(window, s[i-14:i])
            if len(window) == 14:
                return i

print(main())

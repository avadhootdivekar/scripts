import argparse

class converter():
    def __init__(self) -> None:
        self.d = " "
        

    def b2h(self, s: str):
        s = str(s)
        # print(f"input : {s}")
        numList = s.split(self.d)
        h = ""
        for i in numList:
            if len(str(i)) > 0:
                # print(f"i : {i}")
                num = int(i, 2)
                h += f" {hex(num)[2:]}"
        print(f"hex output : {h}")
        return h


    def b2d(self, s: str):
        s = str(s)
        # print(f"input : {s}")
        numList = s.split(self.d)
        d = ""
        for i in numList:
            if len(str(i)) > 0:
                # print(f"i : {i}")
                num = int(i, 2)
                d += f" {int(num)}"
        print(f"decimal output : {d}")
        return d


    def h2b(self, s: str):
        numList = s.split(self.d)
        b = ""
        for i in numList:
            if len(i) > 0:
                num = int(i, 16)
                b += " {0:b}".format(num)
        print(f"binary output : {b}")
        return b


def configParser():
    parser = argparse.ArgumentParser(description="This is default arg parser")
    parser.add_argument(
        "--func", type=str, required=False, help="Function name to be used for parsing")
    parser.add_argument(
        "--s", type=str, required=False, help="String to be parsed")
    parser.add_argument(
        "-d", type=str, required=False, help="Delimeter")
    args = parser.parse_args()
    return args

def main():
    args = configParser()
    print(f"args :  {args}")
    conv = converter()
    conv.d = args.d
    getattr(conv , args.func)(args.s)
    return


main()

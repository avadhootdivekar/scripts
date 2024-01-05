import argparse

class converter():
    def __init__(self) -> None:
        self.d = " "
        self.prefix =  ""
        

    def b2h(self, s: str):
        s = str(s)
        # print(f"input : {s}")
        numList = s.split(self.d)
        h = ""
        for i in numList:
            if len(str(i)) > 0:
                # print(f"i : {i}")
                val = None
                if self.prefix in str(i): 
                    val = str(i).split(self.prefix)
                    if len(val) == 1:
                        val = val[0]
                    else :
                        val = val[1]
                else : 
                    val = str(i)
                num = int(str(val), 2)
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
    parser.add_argument(
        "--prefix", type=str, required=False, help="Prefix of each number")
    args = parser.parse_args()
    return args

def main():
    args = configParser()
    print(f"args :  {args}")
    conv = converter()
    conv.d = args.d
    conv.prefix = args.prefix
    getattr(conv , args.func)(args.s)
    return


main()

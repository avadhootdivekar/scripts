import argparse
import numpy

class converter():
    def __init__(self) -> None:
        self.d = " "
        self.inPrefix =  ""
        self.outDel = ""
        self.outPrefix = ""
        self.inBase = 10
        self.outBase = 10

    def __str__(self) -> str:
        s = f"{self.__dict__}"
        return s
        
    def stripPrefix(self, s:str):
        val = None
        if self.inPrefix in s: 
            val = s.split(self.inPrefix)
            if len(val) == 1:
                val = val[0]
            else :
                val = val[1]
        else : 
            val = s
        return val

    def formatOutput(self, num:list) -> str:
        s = ""
        first=True
        for i in num:
            if not first:
                s += f"{self.outDel}{self.outPrefix}{i}"
            else :
                s += f"{self.outPrefix}{i}"
                first = False

        return s

    def b2h(self, s: str):
        s = str(s)
        # print(f"input : {s}")
        numList = s.split(self.d)
        outNumList = []
        h = ""
        for i in numList:
            if len(str(i)) > 0:
                # print(f"i : {i}")
                val = self.stripPrefix(str(i))
                num = int(str(val), 2)
                outNumList.append(hex(num)[2:])
        h += self.formatOutput(outNumList)
        print(f"hex output : {h}")
        return h


    def b2d(self, s: str):
        s = str(s)
        # print(f"input : {s}")
        numList = s.split(self.d)
        outNumList = []
        d = ""
        for i in numList:
            if len(str(i)) > 0:
                # print(f"i : {i}")
                val = self.stripPrefix(str(i))
                num = int(val, 2)
                outNumList.append(int(num))
        d += self.formatOutput(outNumList)
        print(f"decimal output : {d}")
        return d


    def h2b(self, s: str):
        numList = s.split(self.d)
        b = ""
        outNumList = []
        for i in numList:
            if len(i) > 0:
                val = self.stripPrefix(str(i))
                num = int(val, 16)
                outNumList.append("{0:b}".format(num))
        b += self.formatOutput(outNumList)
        print(f"binary output : {b}")
        return b

    def conv(self, inBase:int , outBase:int , input:str) -> str:
        output = ""
        numList = input.split(self.d)
        intList = []
        outNumList = []
        for i in numList:
            if len(i) > 0:
                val = self.stripPrefix(str(i))
                num = int(val , inBase)
                intList.append(num)
                numStr = numpy.base_repr(num , outBase , padding=0)
                outNumList.append(numStr)
        output += self.formatOutput(outNumList)
        print(f"numbers : {intList} \nOutput : {output}\n")
        return output

def getArg(args, arg:str , default) :
    val = getattr(args , arg)
    if val == None:
        val = default
    return val

def configParser(conv:converter):
    parser = argparse.ArgumentParser(description="This is default arg parser")
    parser.add_argument(
        "--func", type=str, required=False, help="Function name to be used for parsing")
    parser.add_argument(
        "-s" , "--string", type=str, required=False, help="String to be parsed")
    parser.add_argument(
        "-d", "--in-delimeter", type=str, required=False, help="Delimeter")
    parser.add_argument(
        "-p","--in-prefix", type=str, required=False, help="Prefix of each number in the input")
    parser.add_argument(
        "--out-prefix", type=str, required=False, help="Prefix of each number in the output")
    parser.add_argument(
        "--out-delimeter", type=str, required=False, help="Delimeter in the output")
    parser.add_argument(
        "-b","--in-base", type=str, required=False, help="Base of input (0-36)")
    parser.add_argument(
        "--out-base", type=str, required=False, help="Base of output (0-36)")
    parser.add_argument(
        "-i","--in-file", type=str, required=False, help="Base of output (0-36)")
    parser.add_argument(
        "-o","--out-file", type=str, required=False, help="Base of output (0-36)")
    args = parser.parse_args()
    conv.d = getattr(args, "d" , " ")
    conv.inPrefix = getArg(args, "in_prefix" , " ")
    conv.outPrefix = getArg(args, "out_prefix" , " ")
    conv.outDel = getattr(args, "out_delimeter" , "")
    conv.inBase = int(getArg(args, "in_base" , 10))
    conv.outBase = int(getArg(args, "out_base" , 10))
    print( f"argparse : {args} ,\n  conv : {conv} \n" )
    inFile = getArg(args , "in_file", None)
    if inFile !=None : 
        with open(inFile , "r" ) as file:
            args.string = file.read()
            args.string =  args.string.replace('\n' , args.in_delimeter)
    inFile = getArg(args , "in_file", None)
    return args

def main():
    conv = converter()
    args = configParser(conv)
    print(f"args :  {args} , conv : {conv} \n")
    # getattr(conv , args.func)(args.s)
    output = conv.conv(inBase=conv.inBase , outBase=conv.outBase , input=args.string)
    if args.out_file != None : 
        with open(args.out_file , "w+") as file:
            file.write(output)
    return


main()

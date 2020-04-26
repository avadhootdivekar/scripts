import os
import sys
import getopt

ArgList     = ['help',
               'root_dir=',
               'type=',
               'extension=']
#sys.stdout = open('./test.o', 'w')

def List(walk_dir , type , extension , action="Print"):
    ReturnVal   = []
    for root, subdirs, files in os.walk(walk_dir):
    #    filelist = os.listdir(root)
    #    print (filelist)
        if(type=="d"):
            if(action=="Print"):
                print (root)
            elif(action=="Return"):
                ReturnVal.append(root)
        elif(type=="f"):
            for filename in files:
#                 print("Filename : " + filename)
                if filename.endswith(extension):
        #        if "*.c" in files:
        #            sys.stdin.read(1)
                    Line = os.path.join(root , filename)
                    if(action=="Print"):
                        print (Line)
                    elif(action=="Return"):
                        ReturnVal.append(Line)
    return ReturnVal


if __name__ == "__main__":
    root        = None
    type        = None
    extension   = None
    if (len(sys.argv) < 2):
         print("Usage : python ListFiles.py <-r root_dir> <-e file_extension> <-t f,d>")
    optlist,args = getopt.getopt(sys.argv[1:] , 'h' , ArgList )

    for opt in optlist:
#         print("Option : "+ str(opt[0]) + ",   Value : " + opt[1])
        if(opt[0]   == '-h'):
            print("Display help message")
        elif(opt[0]   == '--help'):
            print("Display help message")
        elif(opt[0]     == '--root_dir'):
            root        = opt[1]
        elif(opt[0]     == '--type'):
            type        = opt[1]
        elif(opt[0]     == '--extension'):
            extension   = opt[1]
    if ((root is None) or (type is None) or (extension is None)):
        print("Some parameter missing. Please refer ListFiles.py --help")
    else :
        List(root,type,extension)



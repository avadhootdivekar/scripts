import matplotlib.pyplot as plt
from mpl_toolkits.mplot3d import axes3d
from matplotlib import cm
import numpy as np
import math
import logging


# Creating a logger object
logFormat="[%(asctime)s:%(levelname)s: %(filename)s:%(lineno)s-%(funcName)s() ] %(message)s"
logging.basicConfig(filename="/var/log/sample.log" ,  level=logging.DEBUG , force=True , format=logFormat)
log = logging.getLogger(__name__.split('.')[0])

INVALID_FIELD_VAL	=	0
X_RANGE = [-10 , 10]
Y_RANGE = [-10 , 10]
RESOLUTION = 0.1
POINT_COUNT = int((X_RANGE[1]-X_RANGE[0])/RESOLUTION)
INIT_OBJ_LIST = [
    [-9.5,-9.5],
    [-8,8],
    [8,-9],
    [9,9.5]
]

class obj():
    def __init__(self , position:np.array):
        self.position = position
        return

    def __str__(self):
        s = f"OBJ : position: {self.position}"
        return s

    def getFAt(self, point , default = 0):
        F = INVALID_FIELD_VAL
        d = math.dist(self.position , point)
        if d < 1:
            F = default
        else :
            F = d**(-2)
        log.debug(f"Distance : {d}")
        return F

class P1():
    def __init__(self):
        self.objList = []
        self.X = np.arange(X_RANGE[0] , X_RANGE[1] , RESOLUTION )
        self.Y = np.arange(Y_RANGE[0] , Y_RANGE[1] , RESOLUTION )
        self.X , self.Y = np.meshgrid(self.X , self.Y )
        self.field = np.zeros((POINT_COUNT , POINT_COUNT) )
        return

    def plotField(self):
        ax = plt.figure().add_subplot(projection='3d')
        log.debug(f"field : {self.field}")
        ax.plot_surface(self.X,self.Y,self.field ,  cmap=cm.coolwarm)
        plt.show()
        return

    def initObjList(self):
        for i in range(len(INIT_OBJ_LIST)):
            O1 = obj(np.array(INIT_OBJ_LIST[i]) )
            log.debug(f"Appending obj {O1}")
            self.objList.append(O1)
        log.debug(f"Obj List : {self.objList}")
        return
    
    def calculateField(self):
        defaultVal = 0
        for i in range(len(self.field[0])):
            for j in range(len(self.field)):
                self.field[i][j] = 0
                for eachObj in self.objList:
                    self.field[i][j] += eachObj.getFAt(np.array([self.X[i][j], self.Y[i][j]]) ,default= defaultVal)
                defaultVal = self.field[i][j]
        log.debug(f"Field : {self.field}")
        return

    def sample3dPlot(self):
        ax = plt.figure().add_subplot(projection='3d')
        X, Y, Z = axes3d.get_test_data(0.05)
        ax.contour(X, Y, Z)  # Plot contour curves
        plt.show()
        return

    def sample_2(self):
        ax = plt.figure().add_subplot(projection='3d')
        self.X = np.arange(X_RANGE[0] , X_RANGE[1] , RESOLUTION )
        self.Y = np.arange(Y_RANGE[0] , Y_RANGE[1] , RESOLUTION )
        self.X , self.Y = np.meshgrid(X , Y )

        Z = self.X**2 + self.Y**2
        log.debug(f"X : {self.X} , Y : {self.Y} , Z : {Z}")
        # for i in range(len(X)):
        #     for j in range(len(Y)):
        #         Z[i][j] = i+j
        ax.plot_surface(self.X,self.Y,Z)
        plt.show()
        return
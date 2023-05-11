import matplotlib.pyplot as plt
from mpl_toolkits.mplot3d import axes3d
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

class obj():
    def __init__(self , position:np.array):
        self.position = position
        return

    def getFAt(self, point):
        F = INVALID_FIELD_VAL
        d = math.dist(self.position , point)
        F = d**2
        log.debug(f"Distance : {d}")
        return F

class P1():
    def __init__(self):
        self.objList = []
        return

    def plotField(self):
        return

    def addObj(self, obj):
        self.objList.append(obj)
        return
    
    def calculateField(self):

        return

    def sample3dPlot(self):
        ax = plt.figure().add_subplot(projection='3d')
        X, Y, Z = axes3d.get_test_data(0.05)
        ax.contour(X, Y, Z)  # Plot contour curves
        plt.show()
        return

    def sample_2(self):
        ax = plt.figure().add_subplot(projection='3d')
        X = np.arange(X_RANGE[0] , X_RANGE[1] , RESOLUTION )
        Y = np.arange(Y_RANGE[0] , Y_RANGE[1] , RESOLUTION )
        X , Y = np.meshgrid(X , Y )
        Z = X**2 + Y**2
        log.debug(f"X : {X} , Y : {Y} , Z : {Z}")
        # for i in range(len(X)):
        #     for j in range(len(Y)):
        #         Z[i][j] = i+j
        ax.plot_surface(X,Y,Z)
        plt.show()
        return
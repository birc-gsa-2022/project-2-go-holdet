import pandas as pd
import matplotlib.pyplot as plt
import numpy as np

def makePlot():
    
    file = pd.read_csv("time_search.csv")
    naive4 = np.log(file['naive4'])
    naive8 = np.log(file['naive8'])
    labels = ['naive 4 matches', 'naive 8 matches']

    x_axis = np.log((file['x_size']))
    plt.scatter(x_axis ,naive4, c="blue", label="naive 4 matches, alphabet A*")
    plt.scatter(x_axis ,naive8, c="green", label="naive 8 matches, alphabet A*")
    plt.legend(loc="upper left")
    plt.xlabel(" log (length of pattern)")
    plt.ylabel("log (time)")
    
    #plt.ticklabel_format(useOffset=False, style='plain')
    
    plt.show()



makePlot()
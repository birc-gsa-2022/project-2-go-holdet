import pandas as pd
import matplotlib.pyplot as plt
import numpy as np

def makePlot():
    
    file = pd.read_csv("fixed_n_data.csv")
    
    x_axis = (file['x_size'])
    
    naive10 = file['naive']/x_axis
    naive20 = file['naive_worst']/x_axis
    labels = ['naive 10 matches', 'naive 20 matches']

    
    plt.scatter(x_axis ,naive10, c="blue", label="naive, alphabet A*")
    plt.scatter(x_axis ,naive20, c="green", label="naive, alphabet English*")
    plt.legend(loc="upper left")
    plt.xlabel("length of pattern")
    plt.ylabel("(time ns) / (length of pattern)")
    
    #plt.ticklabel_format(useOffset=False, style='plain')
    
    plt.show()



makePlot()
import matplotlib.pyplot as plt
import pandas as pd
import numpy as np
import sys

def main():
	if len(sys.argv) != 3:
		raise Exception("Usage: viewer.py theta0 theta1")
	df = pd.read_csv("data.csv")

	df.plot(kind = 'scatter', x = 'km', y = 'price')

	theta0 = float(sys.argv[1])
	theta1 = float(sys.argv[2])

	x = np.array(df["km"])
	y = theta0 + theta1 * x

	plt.plot(x, y)
	plt.show()

if __name__ == "__main__":
	try:
		main()
	except Exception as e:
		print(e)

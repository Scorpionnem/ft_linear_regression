#!/bin/bash

python -m venv venv
source venv/bin/activate
pip install pandas
pip install matplotlib
python viewer.py

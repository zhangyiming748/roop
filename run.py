#!/usr/bin/env python3

from roop import core
import os

if __name__ == '__main__':
    os.environ['TK_SILENCE_DEPRECATION'] = '1'
    core.run()

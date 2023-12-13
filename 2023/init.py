import os
import shutil
import subprocess

for i in range(9, 26):
    fld = f"day {i}"
    if os.path.isdir(fld):
        shutil.rmtree(fld)

    shutil.copytree("day 8", fld)

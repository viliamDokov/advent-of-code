from dataclasses import dataclass

all_folders = set()

class Folder:
    def __init__(self, parent, name):
        self.name = name
        self.parent = parent
        self.files = set()
        self.folders = set()

    def __hash__(self):
        return hash(self.name)
    
    def __eq__(self, other):
        return type(other) == type(self) and self.name == other.name
    
    def get_size(self):
        size = 0
        for file in self.files:
            size += file.size
        for folder in self.folders:
            size += folder.get_size()

        all_folders.add(self)
        self.size = size
        return size

class File:
    def __init__(self, size, name):
        self.name = name
        self.size = size
    def __hash__(self):
        return hash(self.name)
    
    def __eq__(self, other):
        return type(other) == type(self) and self.name == other.name



def main():
    with open("inp") as f:
        current = None
        line = f.readline()
        while line.strip() != "":
            # print(line)
            read = False
            command = line[2:].strip()

            if command.startswith("cd"):
                name = command[2:].strip()
                if name != "..":
                    parent  = current
                    current = Folder(parent, name)
                    if parent != None:
                        parent.folders.add(current)
                else:
                    current = current.parent

            elif command.startswith("ls"):
                read = True
                line = f.readline().strip()
                while not line.startswith("$") and line != "":
                    if not line.startswith("dir"):
                        size, name = int(line.split(" ")[0]), line.split(" ")[1]
                        file = File(size,name)
                        print("FILE", file.size, file.name)
                        current.files.add(file)

                    line = f.readline().strip()

            if not read:
                line = f.readline()
    
    while current.parent != None:
        current = current.parent 

    current.get_size()




print(main())
best = 1000000000000000
for folder in all_folders:
    if folder.size >= 3837783:
        if best > folder.size:
            best = folder.size
        print(folder.name, folder.size)

print(best)
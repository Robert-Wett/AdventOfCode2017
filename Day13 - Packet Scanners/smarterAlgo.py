def checkCaught(depth, pos):
    position = pos + 1
    if (position / depth) % 2 == 0:
        idx = depth - (-position % depth) - 1
    else:
        idx = (depth - (position % depth)) - 1
    print idx
    return idx == 0

print checkCaught(3, 0)
//
//  GapBuffer.swift
//  Coder
//
//  Created by yechen on 2020/12/14.
//  Copyright © 2020 Polarnight. All rights reserved.
//

// import Foundation

struct GapBuffer {
    static let initGapSize = 8
    var buffer: [UInt8]
    var gapStart = 0
    var gapSize = Self.initGapSize
    var count: Int {
        buffer.count - gapSize
    }
    
    init() {
        buffer = [UInt8](repeating: 0, count: Self.initGapSize)
    }
    init(initText: String) {
        buffer = [UInt8](initText.utf8)
        buffer.append(contentsOf: [UInt8](repeating: 0, count: Self.initGapSize))
    }
    
    
    
    mutating func insert(index: Int, str: String) {
        guard (0...self.count).contains(index) else {
            print("ERROR >>> insert()")
            return
        }
        guard str.count > 0 else {
            return
        }
        
        if str.count > 1 {
            for (i, c) in str.enumerated() {
                self.insert(index: index+i, str: String(c))
            }
        } else {
            expandGap()
            moveGapToPoint(point: index)
            buffer[gapStart] = UInt8(str.unicodeScalars.first!.value)
            gapStart += 1
            gapSize -= 1
        }
    }
    
    mutating func remove(index: Int) {
        guard (0..<self.count).contains(index) else {
            print("ERROR >>> remove()")
            return
        }
        
        if gapSize == 0 {
            gapStart = index
            gapSize += 1
            return
        }
        
        if (index+1)==gapStart {
            gapStart = index
            gapSize += 1
        } else if index==gapStart {
            gapSize += 1
        } else {
            if index < gapStart {
                moveGapToPoint(point: index+1)
                gapStart = index
                gapSize += 1
            } else {
                moveGapToPoint(point: index)
                gapSize += 1
            }
        }
        
    }
    
    mutating func remove(start: Int, size: Int) {
        var size = size
        if (start+size) > self.count {
            size = self.count - start
        }
        for _ in 0..<size {
            self.remove(index: start)
        }
    }
    
    mutating func expandGap() {
        if gapSize==0 {
            gapStart = buffer.count
            gapSize = Self.initGapSize
            buffer.append(contentsOf: [UInt8](repeating: 0, count: Self.initGapSize))
        }
    }
    
    mutating func moveGapToPoint(point: Int) {
        guard point != gapStart else {
            return
        }
        if point < gapStart {
            // point=0, gapStart=7, gapSize=3
            // abcdefg...
            // abcdef...g
            // abcde...fg
            // ...abcdefg
            for i in 0..<(gapStart-point) {
                buffer[gapStart+gapSize-1-i] = buffer[gapStart-1-i]
            }
        } else if gapStart < point {
            // point=10, gapStart=0, gapSize=3
            // ...abcdefg
            // a...bcdefg
            // ab...cdefg
            // abcdefg...
            for i in 0..<(point-gapStart) {
                buffer[gapStart+i] = buffer[gapStart+gapSize+i]
            }
        }
        gapStart = point
    }
    
    
    
    
    
    func getStatus() -> String {
        var str = ""
        for (index, char) in buffer.enumerated() {
            if gapStart<=index && index<(gapStart+gapSize) {
                str.append(contentsOf: "_")
            } else {
                let c = String( Unicode.Scalar(char) )
                str.append(c)
            }
        }
        return str
    }
    
    func getString() -> String {
        var str = ""
        for (index, char) in buffer.enumerated() {
            if gapStart<=index && index<(gapStart+gapSize) {
            } else {
                let c = String( Unicode.Scalar(char) )
                str.append(c)
            }
        }
        return str
    }
    
    func debug() {
        print("========================================")
        let str = getString()
        print("文字列        :", str)
        print("文字列の長さ   :", str.count)
        print("buffer       :", getStatus())
        print("buffer's size:", buffer.count)
        print("gapStart:", gapStart, "gapSize:", gapSize)
        print()
    }
    
    
}

func testGapBuffer() {
    
    print("\n\n\n===================   新建GapBuffer")
    var gb = GapBuffer()
    gb.debug()
    
    while true {
        print("\n\n\n>>> ", terminator: "")
        if let input = readLine()?.split(separator: " ") {
            let index = Int(input[0])!
            let str = String(input[1])
            print("=================== 往\(index)插入\"\(str)\"")
            gb.insert(index: index, str: str)
            gb.debug()
        } else {
            break
        }
    }
    
    
    
    
}

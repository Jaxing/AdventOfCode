using System;
using System.IO;
using System.Globalization;
using System.Text.RegularExpressions;

public class Start {
    static void Main(string[] args) {
        DateTime start = DateTime.Now;
        string day = args[1];
        string taskNumber = args[2];
        string[] input = File.ReadAllLines(args[3]);

        switch (day) {
            case "1":
                if (taskNumber == "1") {
                    Console.WriteLine(December1.CalculateFrequency(input));
                } else if (taskNumber == "2") {
                    Console.WriteLine(December1.firstRepeatedFrequency(input));
                }
                break;
            case "2":
                if (taskNumber == "1") {
                    Console.WriteLine(December2.checksum(input));
                } else if (taskNumber == "2") {
                    Console.WriteLine(December2.findPackages(input));
                }
                break;
            case "3":
                if (taskNumber == "1") {
                    var overlap = December3.overlap(input);
                    Console.WriteLine(overlap.Count);
                } else if (taskNumber == "2") {
                    Console.WriteLine(December3.firstClaimNotOverlapping(input));
                }
                break;
            case "4":
                if (taskNumber == "1") {
                    Console.WriteLine(December4.sleepyNumber(input));
                } else if (taskNumber == "2") {
                    Console.WriteLine(December4.reliableSleeper(input));
                }
                break;
            case "5":
                if (taskNumber == "1") {
                    Console.WriteLine(December5.unitsLeft(input[0]));
                } else {
                    Console.WriteLine(December5.findBestPolymer(input[0]));
                }
                break;
            case "6":
                if (taskNumber == "1") {
                    Console.WriteLine(December6.largestArea(input));
                } else {
                    Console.WriteLine(December6.sizeOfSafeArea(input));
                }
                break;
            case "7":
                if (taskNumber == "1") {
                    Console.WriteLine(December7.instructionOrder(input));
                } else {
                    Console.WriteLine(December7.timeToFinnish(input));
                }
                break;
            default:
                throw new NotImplementedException();
        }

        Console.WriteLine("Ran in {0} ms", DateTime.Now.Subtract(start).TotalMilliseconds);
    }
}
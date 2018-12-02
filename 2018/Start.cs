using System;
using System.IO;

public class Start {
    static void Main(string[] args) {
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
            default:
                throw new NotImplementedException();
        }
    }
}
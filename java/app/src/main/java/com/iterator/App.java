/*
 * This Java source file was generated by the Gradle 'init' task.
 */
package com.iterator;

import java.time.Instant;
import java.util.ArrayList;
import java.util.List;

public class App {
    private static int NUMBER_OF_ITERATIONS = 1000000000;
    private static int WHEN_PRINT = NUMBER_OF_ITERATIONS - 2;
    private static int NUMBER_OF_REPETITIONS = 10;
    public void doIterations() {
        List<Long> list = new ArrayList<>();
        for (int i = 0; i < NUMBER_OF_REPETITIONS; i++) {
            Instant start = Instant.now();
            for (int j = 0; j < NUMBER_OF_ITERATIONS; j++) {
                int c = j + i;
                if (c == WHEN_PRINT) {
                    Instant end = Instant.now();
                    list.add(end.toEpochMilli() - start.toEpochMilli());
                }
            }
            // Instant end = Instant.now();
            // System.out.println(end.toEpochMilli() - start.toEpochMilli());
        }
        System.out.println(list);
    }

    public static void main(String[] args) {
        new App().doIterations();
    }
}
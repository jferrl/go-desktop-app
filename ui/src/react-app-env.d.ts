/// <reference types="react-scripts" />

declare namespace NodeJS {
  interface Global {
    add: (a: number, b: number) => Promise<number>;
    printObject: (obj: Object) => Promise<void>;
  }
}

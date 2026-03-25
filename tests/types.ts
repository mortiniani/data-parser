// types.ts

export type Data = {
  id: string;
  name: string;
  age: number;
  occupation: string;
};

export type DataArray = Data[];

export type ParsedData = {
  id: string;
  name: string;
  age: number;
  occupation: string;
  parsedOccupation: string;
};

export type ParsedDataArray = ParsedData[];
import React from "react";

import MainLayout from "layouts/MainLayout";

import { Wrap, WrapItem } from "@chakra-ui/react";
import Section from "layouts/Section";

const dummyData = [
  "Item1",
  "Item1",
  "Item1",
  "Item1",
  "Item1",
  "Item1",
  "Item1",
  "Item1",
  "Item1",
];

const HomePage: React.FC<{}> = () => {
  return (
    <MainLayout>
      <Section title={"Today's Webtoon"}>
        <Wrap>
          {dummyData.map((item, idx) => (
            <WrapItem key={idx}>{item}</WrapItem>
          ))}
        </Wrap>
      </Section>
      <Section title={"Recently Update"}>
        <Wrap>
          {dummyData.map((item, idx) => (
            <WrapItem key={idx}>{item}</WrapItem>
          ))}
        </Wrap>
      </Section>
    </MainLayout>
  );
};

export default HomePage;

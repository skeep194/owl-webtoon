import React from "react";

import { Box, Card, CardBody, Heading } from "@chakra-ui/react";

type SectionProps = {
  title: string;
  children: React.ReactElement | React.ReactElement[];
};

const Section: React.FC<SectionProps> = ({ title, children }) => {
  return (
    <Card width={"container.md"} mt={"10px"} mb={"10px"}>
      <CardBody>
        <Heading size={"lg"} mb={"5px"}>
          {title}
        </Heading>
        <Box>{children}</Box>
      </CardBody>
    </Card>
  );
};

export default Section;

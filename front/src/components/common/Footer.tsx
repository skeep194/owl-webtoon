import React from "react";

import { Center, Link, Text } from "@chakra-ui/react";

const Footer: React.FC<{}> = () => {
  return (
    <Center width={"100vw"} height={"5vh"} color={"gray.500"}>
      <Text>
        Present by <Link>BoGwon Kang</Link> & <Link>Jongun Jeong</Link>
      </Text>
    </Center>
  );
};

export default Footer;

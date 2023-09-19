import React from "react";

import LogoImage from "assets/Logo.jpg";
import { Box, Center, HStack, Heading, Image, Link } from "@chakra-ui/react";

const Header: React.FC<{}> = () => {
  return (
    <Center width={"100vw"} height={"100px"}>
      <HStack spacing={5} width={"container.lg"}>
        <Box>
          <Image
            borderRadius={"full"}
            boxSize={"70px"}
            src={LogoImage}
            alt={"owl-webtoon"}
          />
        </Box>
        <Link>
          <Heading size={"md"}>Home</Heading>
        </Link>
        <Link>
          <Heading size={"md"}>Ranking</Heading>
        </Link>
        <Link>
          <Heading size={"md"}>Calendar</Heading>
        </Link>
        <Link>
          <Heading size={"md"}>Community</Heading>
        </Link>
      </HStack>
    </Center>
  );
};

export default Header;

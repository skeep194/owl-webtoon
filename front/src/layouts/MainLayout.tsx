import React from "react";

import { Box, Container } from "@chakra-ui/react";
import Header from "components/common/Header";
import Footer from "components/common/Footer";

type MainLayoutProps = {
  children: React.ReactElement | React.ReactElement[];
};

const MainLayout: React.FC<MainLayoutProps> = ({ children }) => {
  return (
    <Box width={"100vw"} height={"100vh"}>
      <Header />
      <Container
        minH={"calc(95vh - 100px)"}
        maxW={"container.md"}
        centerContent={true}
      >
        {children}
      </Container>
      <Footer />
    </Box>
  );
};
export default MainLayout;

import styled from "styled-components";
import image from "./Image/img.jpeg"


export const ImageBlock = styled.div`
  height: 100vh;
  background-image: url(${image});
  background-repeat: no-repeat;
  background-size: cover;
`;
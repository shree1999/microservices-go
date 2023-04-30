import { useState } from "react";
import Container from "react-bootstrap/Container";
import Row from "react-bootstrap/Row";
import Col from "react-bootstrap/Col";
import axios from "axios";

import CustomButton from "./components/Button";
import Display from "./components/Display";

function App() {
  const [sent, setSent] = useState({});
  const [received, setReceived] = useState({});

  const handleBrokerButton = async () => {
    setSent(() => ({}));
    try {
      const { data } = await axios.get(
        "http://localhost:4000/api/v1/broker/ping"
      );

      setReceived((prevState) => ({ ...prevState, data }));
    } catch (err) {
      console.error(err);
    }
  };

  return (
    <Container>
      <h1 style={{ padding: "0.5rem" }}>
        Test Microservices with React and Golang
      </h1>
      <hr />
      <Row>
        <Col>
          <CustomButton
            text="Test Broker"
            onClickHandler={handleBrokerButton}
            variant="outline-secondary"
          />
        </Col>
      </Row>
      <Display sent={sent} received={received} />
    </Container>
  );
}

export default App;

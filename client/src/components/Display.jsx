import Row from "react-bootstrap/Row";
import Col from "react-bootstrap/Col";

function Display({ sent = {}, received = {} }) {
  return (
    <Row style={{ marginTop: "1rem" }}>
      <Col>
        <h3 style={{ paddingBottom: "0.5rem" }}>Sent</h3>
        <pre style={{ border: "1px solid grey", padding: "1px" }}>
          {JSON.stringify(sent, null, 2)}
        </pre>
      </Col>
      <Col>
        <h3 style={{ paddingBottom: "0.5rem" }}>Received</h3>
        <pre style={{ border: "1px solid grey", padding: "1px" }}>
          {JSON.stringify(received, null, 2)}
        </pre>
      </Col>
    </Row>
  );
}

export default Display;

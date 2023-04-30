import Button from "react-bootstrap/Button";

function CustomButton({ text, onClickHandler, variant }) {
  return (
    <Button variant={variant} onClick={() => onClickHandler()}>
      {text}
    </Button>
  );
}

export default CustomButton;

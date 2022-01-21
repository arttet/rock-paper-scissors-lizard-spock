import { ChoiceItem } from '../Choice';

interface ChoiceProps extends ChoiceItem {
  onClick: (choice: ChoiceItem) => void;
}

const Choice: React.FC<ChoiceProps> = ({ id, name, onClick }) => {
  return (
    <li>
      <button className="button" onClick={() => onClick({id, name})}>{name}</button>
    </li>
  );
};

export default Choice;

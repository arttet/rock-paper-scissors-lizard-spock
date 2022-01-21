import Choice from './Choice';
import Result from './Result';

import { useState } from 'react';
import { ChoiceItems, ChoiceItem } from '../Choice';

import './Game.css';

interface ResultItem {
    results: string;
    player: number;
    computer: number;
}

const Game: React.FC = () => {
    const [baseURL, setBaseURL] = useState('https://codechallenge.boohma.com');
    const [choice, setChoice] = useState<ChoiceItem>();
    const [choices, setChoices] = useState<ChoiceItems>([]);
    const [result, setResult] = useState<string>('');

    const getChoices = () => {
        fetch(baseURL + '/choices', {
            method: 'GET',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json; charset=utf-8',
            },
        })
            .then((response) => response.json())
            .then((data) => {
                console.info('/choices | ', JSON.stringify(data));
                setChoices([...data as ChoiceItems]);
            })
            .catch((error) => {
                setChoices([]);
                console.error('Error:', error);
            });
    }

    const getChoice = () => {
        fetch(baseURL + '/choice', {
            method: 'GET',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json; charset=utf-8',
            },
        })
            .then((response) => response.json())
            .then((data) => {
                console.info('/choice | ', JSON.stringify(data));
                setChoice(data as ChoiceItem)
            })
            .catch((error) => {
                console.error('Error:', error);
            });
    }

    function playWithChoice(player_choice: ChoiceItem) {
        console.log(`played with ${player_choice.id}`);

        fetch(baseURL + '/play', {
            method: 'POST',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json; charset=utf-8',
            },
            body: JSON.stringify({ 'player': player_choice.id })
        })
            .then((response) => response.json())
            .then((data) => {
                console.info('/play | ', JSON.stringify(data));
                const item = data as ResultItem
                setResult(`You played ${choices[item.player - 1].name} & the computer played ${choices[item.computer - 1].name}. You ${item.results}`)
            })
            .catch((error) => {
                console.error('Error:', error);
            });
    }

    return (
        <>
            <h3>Step 1: Put your root URL here</h3>
            <input type="text" value={baseURL} onChange={(e) => setBaseURL(e.target.value)} />

            <h3>Step 2: Populate choices from the /choices endpoint</h3>
            <button className="button" onClick={getChoices}>Click me!</button>

            <ul id="choicesList">
                {
                    choices.map((choice) => (<Choice key={choice.id} id={choice.id} name={choice.name} onClick={playWithChoice} />))
                }
            </ul>

            {choices.length !== 0 ? <Result result={result} /> : null}

            <h2>Random Choice</h2>
            <div id="choice">
                <button className="button" onClick={getChoice}>
                    Get Random Choice from /choice endpoint
                </button>
                <p id="choiceTag">{choice?.name}</p>
            </div>
        </>
    );
};

export default Game;

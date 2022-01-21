interface ResultProps {
    result: string;
}

const Result: React.FC<ResultProps> = ({result}) => {
    return (
        <>
            <h3>Step 3: Click an above choice to play against the computer with the /play endpoint</h3>
            <div id="results">{result}</div>
        </>
    );
};

export default Result;

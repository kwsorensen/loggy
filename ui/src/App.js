// import logo from './logo.svg';
import loggo from './loggy.PNG'
import './App.css';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        {/* <img src={logo} className="App-logo" alt="logo" /> */}
        <img src={loggo} className='App-logo' alt='Loggo the Loggy Logo' />
        <p>
          Logging shit one line at a time. Again.
        </p>
        <a
          className="App-link"
          href="https://github.com/kwsorensen/loggy"
          target="_blank"
          rel="noopener noreferrer"
        >
          Contribute to loggy!
        </a>
      </header>
    </div>
  );
}

export default App;

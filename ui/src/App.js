// import logo from './logo.svg';
import loggo from './loggy.PNG'
import './App.css';
import { Component } from 'react';



class App extends Component {
  constructor(props) {
    super(props)
    this.state = {
      beginning: 5
    }
  }

  componentDidMount() { // calls a function every one second
    this.timerID = setInterval(
      () => this.clickedButton(),
      1000
    );
  }

  componentWillUnmount() {
    clearInterval(this.timerID);
  }

  clickedButton() {
    console.log("Button Clicked")
    this.setState({
      beginning: this.state.beginning + 1
    })
  }

  render() {
  return (
    <div className="App">
      <header className="App-header">
        {/* <img src={logo} className="App-logo" alt="logo" /> */}
        <img src={loggo} className='App-logo' alt='Loggo the Loggy Logo' />
        <p>
          Logging shit one line at a time.
        </p>
        <a
          className="App-link"
          href="https://github.com/kwsorensen/loggy"
          target="_blank"
          rel="noopener noreferrer"
        >
          Contribute to loggy!
        </a>
        <button onClick={() =>this.clickedButton()}>
          Activate Button {this.state.beginning}
        </button>
      </header>
    </div>
  );
}
}



export default App;

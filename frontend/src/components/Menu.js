import React, {Component} from 'react';
import Recipe from "./Recipe";
import "./Menu.css";

class Menu extends Component{
  constructor() {
    super();
    this.state = {
        recipes: []
    }
  }

  componentDidMount() {
    fetch('http://localhost:8080/api/v1/recipes')
        .then( resp => resp.json())
        .then((data)=> {
            this.setState({
                recipes: data.result ? data.result : []
            })
        })
  }

  render() {
    return(
      <article>
        <header>
          <h1>Delicious Recipes</h1>
        </header>
        <div className="recipes">
          {this.state.recipes.map((props, i) => (
            <Recipe key={i} {...props} />
          ))}
        </div>
      </article>
    )
  }

}

export default Menu;

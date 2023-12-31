import json
import questionary
from rich import print
from rich.prompt import Prompt
from rich.console import Console

# Create a console instance for rich text output
console = Console()


def show_about(json_data: str):
    """
    Display the 'About' section from the JSON data.

    Args:
        json_data (str): JSON data as a string.
    """
    data = json.loads(json_data)
    console.print("\n[bold]About[/]:")
    console.print(f"[bright_green]{data['About']}[/]")


def show_contact(json_data: str):
    """
    Display contact information from the JSON data.

    Args:
        json_data (str): JSON data as a string.
    """
    data = json.loads(json_data)
    console.print("\n[bold]Contact me:[/]")
    console.print(f"• [yellow bold]Email[/]: [bold bright_green]{data['Email']}[/]")
    console.print(f"• [yellow bold]Mobile[/]: [bold bright_green]{data['Mobile']}[/]")
    console.print(
        f"• [yellow bold]LinkedIn[/]: [bold bright_magenta]{data['LinkedIn']}[/]"
    )
    console.print(
        f"• [yellow bold]Hackerrank[/]: [bold bright_magenta]{data['Hackerrank']}[/]"
    )
    console.print(
        f"• [yellow bold]Leetcode[/]: [bold bright_magenta]{data['Leetcode']}[/]"
    )
    console.print(
        f"• [yellow bold]GeeksforGeeks[/]: [bold bright_magenta]{data['GeeksforGeeks']}[/]"
    )
    console.print(f"• [yellow bold]Github[/]: [bold bright_magenta]{data['Github']}[/]")
    console.print(
        f"• [yellow bold]Portfolio Website[/]: [bold bright_magenta]{data['Portfolio Website']}[/]"
    )


def show_experience(json_data: str):
    """
    Display work experience details from the JSON data.

    Args:
        json_data (str): JSON data as a string.
    """
    data = json.loads(json_data)
    experiences = data["Experiences"]
    for i, exp in enumerate(experiences, start=1):
        console.print(f"\n[bold]Experience[/] #[bold bright_yellow]{i}[/]")
        console.print(f"[yellow bold]Position[/]: [bright_green]{exp['Position']}[/]")
        console.print(f"[yellow bold]Company[/]: [bright_green]{exp['Company']}[/]")
        console.print(f"[yellow bold]Location[/]: [bright_green]{exp['Location']}[/]")
        console.print(f"[yellow bold]Duration[/]: [bright_magenta]{exp['Duration']}[/]")
        console.print("[yellow bold]Description[/]:")
        for j, desc in enumerate(exp["Description"], start=1):
            console.print(f"• [bright_green]{desc}[/]")


def show_education(json_data: str):
    """
    Display education details from the JSON data.

    Args:
        json_data (str): JSON data as a string.
    """
    data = json.loads(json_data)
    education = data["Education"]
    console.print("\n[bold]Education[/]:")
    for i, entry in enumerate(education, start=1):
        degree = entry["Degree"]
        institution = entry["Institution"]
        duration = entry["Duration"]
        console.print(f"• [bright_yellow]{degree}[/] at [bright_green]{institution}[/]")
        console.print(f"  [bright_magenta]{duration}[/]")


def show_skills(json_data: str):
    """
    Display skills details from the JSON data.

    Args:
        json_data (str): JSON data as a string.
    """
    data = json.loads(json_data)
    console.print("\n[bold]Skills[/]:")
    console.print(f"• [yellow bold]Languages[/]: [bright_green]{data['Languages']}[/]")
    console.print(f"• [yellow bold]Frontend[/]: [bright_green]{data['Frontend']}[/]")
    console.print(f"• [yellow bold]Backend[/]: [bright_green]{data['Backend']}[/]")
    console.print(f"• [yellow bold]Databases[/]: [bright_green]{data['Databases']}[/]")
    console.print(f"• [yellow bold]Cloud[/]: [bright_green]{data['Cloud']}[/]")
    console.print(f"• [yellow bold]AI/ML[/]: [bright_green]{data['AI/ML']}[/]")
    console.print(f"• [yellow bold]Workspace[/]: [bright_green]{data['Workspace']}[/]")


def show_projects(json_data: str):
    """
    Display project details from the JSON data.

    Args:
        json_data (str): JSON data as a string.
    """
    data = json.loads(json_data)
    projects = data["Projects"]
    for i, project in enumerate(projects, start=1):
        console.print(f"\n[bold]Project[/] #[bold bright_yellow]{i}[/]")
        console.print(f"[bold bright_yellow]{project['Title']}[/]")
        console.print("[yellow bold]Description[/]:")
        for j, desc in enumerate(project["Description"], start=1):
            console.print(f"• [bright_green]{desc}[/]")
        console.print(
            f"[yellow bold]Source Code[/]: [bright_magenta]{project['Source Code']}[/]"
        )
        console.print(
            f"[yellow bold]Application Link[/]: [bright_magenta]{project['Application Link']}[/]"
        )


def show_certificates(json_data: str):
    """
    Display certificate details from the JSON data.

    Args:
        json_data (str): JSON data as a string.
    """
    data = json.loads(json_data)
    certificates = data["Certificates"]
    console.print("\n[bold]Certificates[/]:")
    for i, certificate in enumerate(certificates, start=1):
        console.print(
            f"• [bright_green]{certificate.split('-')[0]}[/]-[bold bright_magenta]{certificate.split('-')[1]}[/]"
        )


def main():
    console.print(
        "\nHey there! I'm [bold bright_yellow]Siva Prakash[/], I'm a [bold bright_green]machine learning enthusiast[/] from India and now currently working with some self paced mini-projects in the field of data analytics, software engineering and machine learning."
    )

    options = [
        "About",
        "Experience",
        "Education",
        "Skills",
        "Projects",
        "Certificates",
        "Contact",
        "Exit",
    ]

    while True:
        choice = questionary.select(
            "\nWhat would you like to know?", choices=options, qmark=""
        ).ask()

        if choice == options[0]:
            file_path = "./data/about/about.json"
            try:
                with open(file_path, "r") as file:
                    contents = file.read()
                    show_about(contents)
            except FileNotFoundError:
                console.print("Couldn't find or load that file.")

        elif choice == options[1]:
            file_path = "./data/experience/experience.json"
            try:
                with open(file_path, "r") as file:
                    contents = file.read()
                    show_experience(contents)
            except FileNotFoundError:
                console.print("Couldn't find or load that file.")

        elif choice == options[2]:
            file_path = "./data/education/education.json"
            try:
                with open(file_path, "r") as file:
                    contents = file.read()
                    show_education(contents)
            except FileNotFoundError:
                console.print("Couldn't find or load that file.")

        elif choice == options[3]:
            file_path = "./data/skills/skills.json"
            try:
                with open(file_path, "r") as file:
                    contents = file.read()
                    show_skills(contents)
            except FileNotFoundError:
                console.print("Couldn't find or load that file.")

        elif choice == options[4]:
            file_path = "./data/projects/projects.json"
            try:
                with open(file_path, "r") as file:
                    contents = file.read()
                    show_projects(contents)
            except FileNotFoundError:
                console.print("Couldn't find or load that file.")

        elif choice == options[5]:
            file_path = "./data/certificates/certificates.json"
            try:
                with open(file_path, "r") as file:
                    contents = file.read()
                    show_certificates(contents)
            except FileNotFoundError:
                console.print("Couldn't find or load that file.")

        elif choice == options[6]:
            file_path = "./data/contact/contact.json"
            try:
                with open(file_path, "r") as file:
                    contents = file.read()
                    show_contact(contents)
            except FileNotFoundError:
                console.print("Couldn't find or load that file.")

        else:
            console.print("[bold bright_yellow] Bye! Have a great day!\n[/]")
            break


if __name__ == "__main__":
    main()

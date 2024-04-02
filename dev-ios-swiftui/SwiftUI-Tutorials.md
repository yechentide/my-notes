# SwiftUI Tutorials

## Creating and Combining Views

ライブラリを見れば、どんな部品(View)があるのかがわかる
```swift
Text() Spacer() Image()
VStack() HStack()

Text("Turtle Rock")
    .font(.title)
    .foregroundColor(.green)
    .padding()

Image("turtlerock")
    .clipShape(Circle())
    .overlay {
        Circle().stroke(.white, lineWidth: 4)
    }
    .shadow(radius: 7)

CircleImage()
    .offset(y: -130)
    .padding(.bottom, -130)

.ignoresSafeArea(edges: .top)
```

## Building Lists and Navigation

Adding `Codable` conformance makes it easier to move data between the structure and a data file.
You’ll rely on the Decodable component of the `Codable` protocol later in this section to read data from file.
```swift
struct Landmark: Hashable, Codable {}

func load<T: Decodable>(_ filename: String) -> T {
    let data: Data

    guard let file = Bundle.main.url(forResource: filename, withExtension: nil) else {
        fatalError("Couldn't find \(filename) in main bundle.")
    }

    do {
        data = try Data(contentsOf: file)
    } catch {
        fatalError("Couldn't load \(filename) from main bundle:\n\(error)")
    }

    do {
        let decoder = JSONDecoder()
        return try decoder.decode(T.self, from: data)
    } catch {
        fatalError("Couldn't parse \(filename) as \(T.self):\n\(error)")
    }
}

Image()
    .resizable()
    .frame(width: 50, height: 50)

Group {
    // 複数のプレビュー
}
.previewLayout(.fixed(width: 300, height: 70))
.previewDevice(PreviewDevice(rawValue: "iPhone SE (2nd generation)"))

ForEach(["iPhone SE (2nd generation)", "iPhone XS Max"], id: \.self) { deviceName in
    LandmarkList()
        .previewDevice(PreviewDevice(rawValue: deviceName))
        .previewDisplayName(deviceName)
}

List{}
List(landmarks, id: \.id) { landmark in }   // or conform Identifiable protocol
List {
      ForEach() {}
      View()
}

NavigationView {
    NavigationLink { /*遷移先*/ } label: { /*表示するもの*/ }
    .navigationTitle("Landmarks")
    .navigationBarTitleDisplayMode(.inline)
}

ScrollView {}

Toggle(isOn: $showFavoritesOnly) { Text("Favorites only") }

final class ModelData: ObservableObject {
    @Published var landmarks: [Landmark] = load("landmarkData.json")
}
struct LandmarkList: View {
    @EnvironmentObject var modelData: ModelData
}
static var previews: some View {
    LandmarkList()
        .environmentObject(ModelData())
}
@main
struct LandmarksApp: App {
    @StateObject private var modelData = ModelData()
    var body: some Scene {
        WindowGroup {
            ContentView()
                .environmentObject(modelData)
        }
    }
}

Label("Toggle Favorite", systemImage: "star.fill")
    .labelStyle(.iconOnly)
```

## Drawing Paths and Shapes

```swift
Path { path in
    var width: CGFloat = 100.0
    let height = width
    path.move(to: CGPoint(x: width * 0.95, y: height * 0.20))
}
.fill(.black)
.stroke()
.strokeBorder()
```

## Animating Views and Transitions

```swift
Label("Graph", systemImage: "chevron.right.circle")
    .labelStyle(.iconOnly)
    .imageScale(.large)
    .rotationEffect(.degrees(showDetail ? 90 : 0))   // <--
    .animation(nil, value: showDetail)               // <--
    .scaleEffect(showDetail ? 1.5 : 1)               // <--
    .padding()
    .animation(.easeInOut, value: showDetail)        // <--
    // .easeInOut   .spring()

Button {
    withAnimation {
        showDetail.toggle()
    }
    withAnimation(.easeInOut(duration: 4)) {
        showDetail.toggle()
    }
} label: {
    Text("animete")
}

View().transition(.slide)

extension AnyTransition {
    static var moveAndFade: AnyTransition {
        AnyTransition.move(edge: .trailing)
    }
    static var moveAndFade: AnyTransition {
        .asymmetric(
            insertion: .move(edge: .trailing).combined(with: .opacity),
            removal: .scale.combined(with: .opacity)
        )
    }
    static func ripple() -> Animation {
        Animation.spring(dampingFraction: 0.5)
            .speed(2)
            .delay(0.03)
    }
}
View().transition(.moveAndFade)
```

## Composing Complex Interfaces

```swift
Dictionary(
    grouping: landmarks,
    by: { $0.category.rawValue }
)

List {
    Image()
        .renderingMode(.original)
        .resizable()
        .scaledToFill()
        .frame(height: 200)
        .clipped()
        .listRowInsets(EdgeInsets())
}

struct ContentView: View {
    @State private var selection: Tab = .featured

    enum Tab {
        case featured
        case list
    }

    var body: some View {
        TabView(selection: $selection) {
            CategoryHome()
                .tabItem {
                    Label("Featured", systemImage: "star")
                }
                .tag(Tab.featured)

            LandmarkList()
                .tabItem {
                    Label("List", systemImage: "list.bullet")
                }
                .tag(Tab.list)
        }
    }
}
```

## Working with UI Controls

```swift
@Environment(\.editMode) var editMode
.accessibilityLabel(_:)

List{}.listStyle(.inset)
EditButton()

List {
    Toggle(isOn: $profile.prefersNotifications) {
        Text("Enable Notifications").bold()
    }
    Picker("Seasonal Photo", selection: $profile.seasonalPhoto) {
        ForEach(Profile.Season.allCases) { season in
            Text(season.rawValue).tag(season)
        }
    }
    .pickerStyle(.segmented)
}

if editMode?.wrappedValue == .active {
    Button("Cancel", role: .cancel) {
        draftProfile = modelData.profile
        editMode?.animation().wrappedValue = .inactive
    }
}
```

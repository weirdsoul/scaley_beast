x=[0,35,50,65,80,130,170,190]
y=[0,45,90,135,180,270,330,360]
p=polyfit(x,y,4)
polyout(p,'x')

xx=0:1:190
f=polyval(p,xx)


plot(x,y,'o',xx,f,'-')
ylabel ("Degrees");
xlabel ("Speed (kts)");

